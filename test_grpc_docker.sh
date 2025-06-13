#!/bin/bash

echo "=== TESTANDO gRPC NO DOCKER ==="
echo ""

# Verificar se grpcurl está instalado
if ! command -v grpcurl &> /dev/null; then
    echo "📦 Instalando grpcurl..."
    
    # Para Ubuntu/Debian
    if command -v apt &> /dev/null; then
        sudo apt update && sudo apt install -y grpcurl
    # Para sistemas com Go instalado
    elif command -v go &> /dev/null; then
        go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
        export PATH=$PATH:$(go env GOPATH)/bin
    else
        echo "❌ Não foi possível instalar grpcurl automaticamente"
        echo "💡 Instale manualmente:"
        echo "   Ubuntu/Debian: sudo apt install grpcurl"
        echo "   Go: go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest"
        exit 1
    fi
fi

echo "✅ grpcurl instalado: $(grpcurl --version)"
echo ""

# Testar conexão gRPC
echo "🔍 Listando serviços gRPC disponíveis:"
grpcurl -plaintext localhost:50051 list

echo ""
echo "📋 Listando métodos do OrderService:"
grpcurl -plaintext localhost:50051 list pb.OrderService

echo ""
echo "📝 Testando ListOrders:"
grpcurl -plaintext localhost:50051 pb.OrderService/ListOrders

echo ""
echo "➕ Testando CreateOrder:"
grpcurl -plaintext \
  -d '{
    "id": "grpc-docker-test-001", 
    "price": 150.0, 
    "tax": 15.0
  }' \
  localhost:50051 pb.OrderService/CreateOrder

echo ""
echo "📋 Listando pedidos novamente para verificar criação:"
grpcurl -plaintext localhost:50051 pb.OrderService/ListOrders

echo ""
echo "✅ Teste gRPC concluído!" 