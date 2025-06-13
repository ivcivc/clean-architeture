# 🧪 **Guia Completo de Testes - gRPC e GraphQL**

## **🔧 Testando gRPC**

### **✅ gRPC está FUNCIONANDO perfeitamente!**

**Comandos para testar:**

```bash
# 1. Listar serviços disponíveis
./grpcurl -plaintext localhost:50051 list

# 2. Listar métodos do OrderService
./grpcurl -plaintext localhost:50051 list pb.OrderService

# 3. Testar ListOrders (listar todas as orders)
./grpcurl -plaintext localhost:50051 pb.OrderService.ListOrders

# 4. Testar CreateOrder (criar nova order)
./grpcurl -plaintext -d '{"id": "grpc-test-999", "price": 500, "tax": 50}' localhost:50051 pb.OrderService.CreateOrder
```

**Resultados obtidos:**
- ✅ **ListOrders**: Retorna todas as orders do banco
- ✅ **CreateOrder**: Cria nova order e calcula final_price automaticamente
- ✅ **Porta 50051**: Funcionando corretamente

---

## **🔧 Testando GraphQL**

### **⚠️ GraphQL com problema de inicialização**

**Problema identificado:** O container está crashando devido ao RabbitMQ não estar pronto no momento da inicialização.

**Soluções:**

### **Opção 1: Testar via Playground (Recomendado)**
```bash
# Acesse no navegador:
http://localhost:8081/

# Query para testar:
{
  listOrders {
    id
    price
    tax
    finalPrice
  }
}

# Mutation para criar order:
mutation {
  createOrder(input: {
    id: "graphql-test-123"
    price: 400
    tax: 40
  }) {
    id
    price
    tax
    finalPrice
  }
}
```

### **Opção 2: Testar via cURL**
```bash
# ListOrders
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"query": "{ listOrders { id price tax finalPrice } }"}' \
  http://localhost:8081/query

# CreateOrder
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"query": "mutation { createOrder(input: {id: \"gql-test-456\", price: 600, tax: 60}) { id price tax finalPrice } }"}' \
  http://localhost:8081/query
```

---

## **🐛 Correção do Problema RabbitMQ**

**Problema:** Container crasha porque RabbitMQ não está pronto.

**Solução aplicada:**
1. ✅ Aumentei delay de 2s → 5s entre tentativas
2. ✅ Mantive 10 tentativas de reconexão
3. ✅ Adicionei logs detalhados

**Para aplicar a correção:**
```bash
# Recompilar com as correções
docker-compose down
docker-compose up --build -d

# Aguardar inicialização (30-60 segundos)
sleep 60

# Testar GraphQL
curl -X POST -H "Content-Type: application/json" -d '{"query": "{ listOrders { id price tax finalPrice } }"}' http://localhost:8081/query
```

---

## **📊 Status Atual dos Serviços**

| Serviço | Porta | Status | Funcionalidade |
|---------|-------|--------|----------------|
| **REST API** | 8080 | ✅ **FUNCIONANDO** | GET/POST /order |
| **gRPC** | 50051 | ✅ **FUNCIONANDO** | CreateOrder, ListOrders |
| **GraphQL** | 8081 | ⚠️ **PROBLEMA INIT** | Playground acessível, queries com erro |
| **MySQL** | 3308 | ✅ **FUNCIONANDO** | Dados persistindo |
| **RabbitMQ** | 15672 | ✅ **FUNCIONANDO** | Management UI ativo |

---

## **🎯 Próximos Passos**

1. **gRPC**: ✅ **COMPLETO** - Todos os testes passando
2. **GraphQL**: 🔧 **EM CORREÇÃO** - Aplicar fix do RabbitMQ
3. **Documentação**: ✅ **COMPLETA** - README.md atualizado
4. **Testes**: ✅ **FUNCIONAIS** - api.http com exemplos

---

## **💡 Dicas de Uso**

### **Para gRPC:**
- Use `grpcurl` para testes rápidos
- Reflection está habilitado (lista métodos automaticamente)
- Suporte completo a CreateOrder e ListOrders

### **Para GraphQL:**
- Playground em http://localhost:8081/ é mais fácil para testes
- Schema completo com queries e mutations
- Integração com mesmo banco de dados

### **Para REST:**
- Use arquivo `api.http` para testes
- Endpoints: GET/POST /order
- Mesma funcionalidade dos outros protocolos 