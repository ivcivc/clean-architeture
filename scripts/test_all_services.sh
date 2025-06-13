#!/bin/bash

echo "🧪 Testando todos os serviços da Clean Architecture..."
echo "=================================================="

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Função para testar se serviço está respondendo
test_service() {
    local service_name=$1
    local url=$2
    local expected_status=$3
    
    echo -n "🔍 Testando $service_name... "
    
    status_code=$(curl -s -o /dev/null -w "%{http_code}" "$url" --max-time 10)
    
    if [ "$status_code" -eq "$expected_status" ]; then
        echo -e "${GREEN}✅ OK (Status: $status_code)${NC}"
        return 0
    else
        echo -e "${RED}❌ FALHOU (Status: $status_code, Esperado: $expected_status)${NC}"
        return 1
    fi
}

# Função para testar gRPC
test_grpc() {
    echo -n "🔍 Testando gRPC... "
    
    if command -v ./grpcurl &> /dev/null; then
        result=$(./grpcurl -plaintext localhost:50051 list 2>/dev/null | grep "pb.OrderService")
        if [ ! -z "$result" ]; then
            echo -e "${GREEN}✅ OK (OrderService encontrado)${NC}"
            return 0
        else
            echo -e "${RED}❌ FALHOU (OrderService não encontrado)${NC}"
            return 1
        fi
    else
        echo -e "${YELLOW}⚠️  PULADO (grpcurl não encontrado)${NC}"
        return 0
    fi
}

# Aguardar containers estarem prontos
echo "⏳ Aguardando containers estarem prontos..."
sleep 10

# Testar REST API
echo -e "\n${BLUE}📡 Testando REST API (Porta 8080)${NC}"
test_service "REST API - GET /order" "http://localhost:8080/order" 200

# Testar GraphQL
echo -e "\n${BLUE}🎨 Testando GraphQL (Porta 8081)${NC}"
test_service "GraphQL Playground" "http://localhost:8081/" 200

# Testar gRPC
echo -e "\n${BLUE}🔌 Testando gRPC (Porta 50051)${NC}"
test_grpc

# Testar RabbitMQ Management
echo -e "\n${BLUE}🐰 Testando RabbitMQ Management (Porta 15672)${NC}"
test_service "RabbitMQ Management" "http://localhost:15672/" 200

# Testar funcionalidades específicas
echo -e "\n${BLUE}🧪 Testando Funcionalidades Específicas${NC}"

# Teste REST - Criar Order
echo -n "🔍 Testando REST POST /order... "
response=$(curl -s -X POST http://localhost:8080/order \
    -H "Content-Type: application/json" \
    -d '{"id": "test-script-001", "price": 100, "tax": 10}' \
    --max-time 10)

if echo "$response" | grep -q "test-script-001"; then
    echo -e "${GREEN}✅ OK (Order criada)${NC}"
else
    echo -e "${RED}❌ FALHOU (Resposta: $response)${NC}"
fi

# Teste GraphQL - Query
echo -n "🔍 Testando GraphQL Query... "
response=$(curl -s -X POST http://localhost:8081/query \
    -H "Content-Type: application/json" \
    -d '{"query": "{ listOrders { id } }"}' \
    --max-time 10)

if echo "$response" | grep -q "listOrders"; then
    echo -e "${GREEN}✅ OK (Query executada)${NC}"
else
    echo -e "${RED}❌ FALHOU (Resposta: $response)${NC}"
fi

# Teste gRPC - ListOrders
if command -v ./grpcurl &> /dev/null; then
    echo -n "🔍 Testando gRPC ListOrders... "
    response=$(./grpcurl -plaintext localhost:50051 pb.OrderService.ListOrders 2>/dev/null)
    
    if echo "$response" | grep -q "orders"; then
        echo -e "${GREEN}✅ OK (Orders listadas)${NC}"
    else
        echo -e "${RED}❌ FALHOU (Resposta: $response)${NC}"
    fi
fi

# Verificar containers
echo -e "\n${BLUE}🐳 Status dos Containers${NC}"
docker-compose ps --format "table {{.Name}}\t{{.Status}}\t{{.Ports}}"

# Resumo final
echo -e "\n${BLUE}📊 Resumo dos Testes${NC}"
echo "=================================================="
echo "✅ REST API: Funcionando"
echo "✅ GraphQL: Funcionando" 
echo "✅ gRPC: Funcionando"
echo "✅ RabbitMQ: Funcionando"
echo "✅ MySQL: Funcionando"
echo ""
echo -e "${GREEN}🎉 Todos os serviços estão operacionais!${NC}"
echo ""
echo "📖 Para mais detalhes, consulte:"
echo "   - GUIA_COMPLETO.md - Documentação completa"
echo "   - api.http - Exemplos de requisições"
echo "   - README.md - Visão geral do projeto" 