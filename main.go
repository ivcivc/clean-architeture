package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"time"

	gqlgen_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/devfullcycle/20-CleanArch/configs"
	"github.com/devfullcycle/20-CleanArch/internal/event/handler"
	"github.com/devfullcycle/20-CleanArch/internal/infra/database"
	"github.com/devfullcycle/20-CleanArch/internal/infra/graph"
	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/service"
	"github.com/devfullcycle/20-CleanArch/internal/infra/web/webserver"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrdersUseCase := NewListOrdersUseCase(db)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order", "POST", webOrderHandler.Create)
	webserver.AddHandler("/order", "GET", webOrderHandler.List)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	// GraphQL
	srv := gqlgen_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrdersUseCase:  *listOrdersUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	go http.ListenAndServe(":"+configs.GraphQLServerPort, nil)

	fmt.Println("All services are running:")
	fmt.Println("REST API: http://localhost:8080/order")
	fmt.Println("gRPC: localhost:50051")
	fmt.Println("GraphQL: http://localhost:8081")

	// Keep the main goroutine alive
	select {}
}

func getRabbitMQChannel() *amqp.Channel {
	var conn *amqp.Connection
	var err error

	// Retry connection up to 10 times with 5 second delay
	for i := 0; i < 10; i++ {
		conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err == nil {
			break
		}
		fmt.Printf("Failed to connect to RabbitMQ (attempt %d/10): %v\n", i+1, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to RabbitMQ after 10 attempts: %v", err))
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}

func NewListOrdersUseCase(db *sql.DB) *usecase.ListOrdersUseCase {
	orderRepository := database.NewOrderRepository(db)
	return usecase.NewListOrdersUseCase(orderRepository)
}
