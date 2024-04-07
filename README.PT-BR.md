# Guia de Execução do Desafio

Este guia detalha os passos necessários para preparar e executar os componentes do desafio de clean architecture. Siga as instruções abaixo para testar as APIs REST e GraphQL, além da API gRPC.

## Preparação do Ambiente

Antes de iniciar os testes, é necessário preparar o ambiente com os seguintes passos:

1. **Iniciar os Serviços com Docker Compose**

    Execute o comando abaixo para iniciar todos os serviços necessários usando Docker Compose. Este comando deve ser rodado na raiz do projeto:

    ```bash
    docker compose up -d
    ```

## Testando as APIs

### API REST

Para testar a API REST, utilize os arquivos HTTP disponíveis no diretório `/api`. Siga as instruções abaixo:

- **Criar Pedido**: Use o arquivo `/api/create_order.http` para criar um novo pedido.
- **Listar Pedidos**: Use o arquivo `/api/list_order.http` para listar todos os pedidos existentes.

### API GraphQL

Para testar a API GraphQL, siga os passos abaixo:

1. Copie o conteúdo do arquivo `/api/mutations.graph`.
2. Acesse a interface gráfica do GraphQL em [http://localhost:8080/manager/html/](http://localhost:8080/manager/html/).
3. Cole o conteúdo copiado na interface gráfica e execute a consulta ou mutação.

### API gRPC

Para testar a API gRPC, utilize a ferramenta Evans. Execute o comando abaixo na raiz do projeto:

```bash
evans --host 127.0.0.1 ./internal/infra/grpc/protofiles/order.proto
```

### Verificando Registros no RabbitMQ

- Para verificar os registros no RabbitMQ, acesse o seguinte endereço no seu navegador:

  [http://localhost:15672/#/queues](http://localhost:15672/#/queues). Este link levará à interface do RabbitMQ, onde você pode verificar as filas e mensagens.
