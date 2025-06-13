#!/bin/bash

echo "🚀 Iniciando Sistema Clean Architecture"
echo "======================================"

# Parar containers existentes
echo "🛑 Parando containers existentes..."
docker-compose down

# Limpar volumes se necessário
echo "🧹 Limpando volumes antigos..."
docker-compose down -v

# Construir e iniciar containers
echo "🔨 Construindo e iniciando containers..."
docker-compose up --build -d

# Aguardar containers iniciarem
echo "⏳ Aguardando containers iniciarem..."
sleep 30

# Verificar status
echo "📊 Status dos containers:"
docker-compose ps

# Executar migrações manualmente
echo "📊 Executando migrações do banco..."
docker exec mysql mysql -u root -proot -e "CREATE DATABASE IF NOT EXISTS orders;"
docker exec -i mysql mysql -u root -proot orders < sql/migrations/001_create_orders_table.sql

echo "✅ Sistema iniciado com sucesso!"
echo ""
echo "🔗 URLs dos serviços:"
echo "📡 REST API: http://localhost:8080/order"
echo "🎨 GraphQL: http://localhost:8081/"
echo "🔌 gRPC: localhost:50051"
echo "🐰 RabbitMQ: http://localhost:15672/ (guest/guest)"
echo ""
echo "🧪 Para testar, execute:"
echo "curl http://localhost:8080/order" 