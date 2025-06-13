# 🏗️ Clean Architecture Go - Sistema de Pedidos

Um projeto completo implementando **Clean Architecture** em Go com múltiplos protocolos de comunicação: **REST API**, **gRPC** e **GraphQL**.

## 🚀 Características

- ✅ **Clean Architecture** - Separação clara de responsabilidades
- ✅ **REST API** - Endpoints HTTP tradicionais
- ✅ **gRPC** - Comunicação de alta performance
- ✅ **GraphQL** - API flexível com queries customizáveis
- ✅ **MySQL** - Persistência de dados
- ✅ **RabbitMQ** - Mensageria assíncrona
- ✅ **Docker** - Containerização completa
- ✅ **Event-Driven** - Arquitetura orientada a eventos

## 🏛️ Arquitetura

```
├── cmd/ordersystem/          # Ponto de entrada da aplicação
├── configs/                  # Configurações
├── internal/
│   ├── entity/              # Entidades de negócio
│   ├── usecase/             # Casos de uso
│   ├── infra/
│   │   ├── database/        # Repositórios MySQL
│   │   ├── web/             # Handlers REST
│   │   ├── grpc/            # Serviços gRPC
│   │   └── graph/           # Resolvers GraphQL
│   └── event/               # Handlers de eventos
├── pkg/events/              # Sistema de eventos
├── sql/                     # Scripts SQL
└── docker-compose.yaml      # Orquestração Docker
```

## 🛠️ Tecnologias

- **Go 1.23** - Linguagem principal
- **MySQL 5.7** - Banco de dados
- **RabbitMQ** - Message broker
- **gRPC** - Comunicação RPC
- **GraphQL** - API query language
- **Docker & Docker Compose** - Containerização

## 🚀 Como Executar

### Pré-requisitos
- Docker
- Docker Compose

### Execução
```bash
# Clonar o repositório
git clone git@github.com:ivcivc/clean-architeture.git
cd clean-architeture

# Iniciar todos os serviços
docker-compose up -d

# Verificar status
docker-compose ps
```

## 📡 Endpoints Disponíveis

### REST API (Porta 8080)
```bash
# Listar pedidos
curl http://localhost:8080/order

# Criar pedido
curl -X POST http://localhost:8080/order \
  -H "Content-Type: application/json" \
  -d '{"id": "123", "price": 100.0, "tax": 10.0}'
```

### gRPC (Porta 50051)
```bash
# Instalar grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# Listar pedidos
grpcurl -plaintext localhost:50051 pb.OrderService/ListOrders

# Criar pedido
grpcurl -plaintext -d '{"id": "123", "price": 100.0, "tax": 10.0}' \
  localhost:50051 pb.OrderService/CreateOrder
```

### GraphQL (Porta 8081)
Acesse o **GraphQL Playground**: http://localhost:8081

**Query de exemplo:**
```graphql
query {
  listOrders {
    id
    price
    tax
    finalPrice
  }
}
```

**Mutation de exemplo:**
```graphql
mutation {
  createOrder(input: {
    id: "123"
    price: 100.0
    tax: 10.0
  }) {
    id
    price
    tax
    finalPrice
  }
}
```

## 🗄️ Banco de Dados

**MySQL** rodando na porta **3308**:
```bash
# Conectar ao MySQL
mysql -h localhost -P 3308 -u root -p
# Senha: root
```

## 📨 RabbitMQ

**Management UI** disponível em: http://localhost:15672
- **Usuário:** guest
- **Senha:** guest

## 🧪 Testes

```bash
# Testar REST API
curl http://localhost:8080/order

# Testar gRPC
grpcurl -plaintext localhost:50051 list

# Testar GraphQL
curl -X POST http://localhost:8081/query \
  -H "Content-Type: application/json" \
  -d '{"query": "{ listOrders { id price tax finalPrice } }"}'
```

## 📁 Estrutura do Projeto

### Entidades
- **Order** - Entidade principal do sistema

### Casos de Uso
- **CreateOrderUseCase** - Criar novo pedido
- **ListOrdersUseCase** - Listar pedidos

### Infraestrutura
- **OrderRepository** - Persistência MySQL
- **WebOrderHandler** - Handlers REST
- **OrderService** - Serviços gRPC
- **Resolver** - Resolvers GraphQL

### Eventos
- **OrderCreated** - Evento disparado na criação de pedidos
- **OrderCreatedHandler** - Handler que envia para RabbitMQ

## 🐳 Docker

O projeto inclui:
- **Dockerfile** multi-stage para a aplicação Go
- **docker-compose.yaml** com todos os serviços
- **Health checks** para garantir inicialização correta

## 🔧 Configuração

As configurações são carregadas via variáveis de ambiente:
- `DB_HOST` - Host do MySQL
- `DB_PORT` - Porta do MySQL
- `DB_USER` - Usuário do MySQL
- `DB_PASSWORD` - Senha do MySQL
- `DB_NAME` - Nome do banco
- `WEB_SERVER_PORT` - Porta do servidor web
- `GRPC_SERVER_PORT` - Porta do servidor gRPC
- `GRAPHQL_SERVER_PORT` - Porta do servidor GraphQL

## 📈 Status dos Serviços

| Serviço | Porta | Status | Endpoint |
|---------|-------|--------|----------|
| REST API | 8080 | ✅ | http://localhost:8080/order |
| gRPC | 50051 | ✅ | localhost:50051 |
| GraphQL | 8081 | ✅ | http://localhost:8081 |
| MySQL | 3308 | ✅ | localhost:3308 |
| RabbitMQ | 15672 | ✅ | http://localhost:15672 |

## 🤝 Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature
3. Commit suas mudanças
4. Push para a branch
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT.

---

**Desenvolvido com ❤️ usando Clean Architecture em Go** 