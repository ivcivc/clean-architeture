#!/bin/bash

echo "🚀 Iniciando aplicação Clean Architecture..."

# Aguardar MySQL estar pronto
echo "⏳ Aguardando MySQL estar pronto..."
while ! mysqladmin ping -h"mysql" -P"3306" -u"root" -p"root" --silent; do
    echo "MySQL não está pronto ainda. Aguardando..."
    sleep 2
done
echo "✅ MySQL está pronto!"

# Aguardar RabbitMQ estar pronto
echo "⏳ Aguardando RabbitMQ estar pronto..."
while ! nc -z rabbitmq 5672; do
    echo "RabbitMQ não está pronto ainda. Aguardando..."
    sleep 2
done
echo "✅ RabbitMQ está pronto!"

# Executar migrações
echo "📊 Executando migrações do banco de dados..."
mysql -h mysql -P 3306 -u root -proot orders < /app/sql/migrations/001_create_orders_table.sql
echo "✅ Migrações executadas com sucesso!"

# Inserir dados de exemplo se não existirem
echo "📝 Inserindo dados de exemplo..."
mysql -h mysql -P 3306 -u root -proot orders -e "
INSERT IGNORE INTO orders (id, price, tax, final_price) VALUES 
('example-001', 100.0, 10.0, 110.0),
('example-002', 200.0, 20.0, 220.0),
('example-003', 300.0, 30.0, 330.0);
"
echo "✅ Dados de exemplo inseridos!"

echo "🎯 Iniciando serviços..."
echo "📡 REST API será iniciada na porta 8080"
echo "🔌 gRPC será iniciado na porta 50051"  
echo "🎨 GraphQL será iniciado na porta 8081"

# Iniciar a aplicação
exec /app/main 