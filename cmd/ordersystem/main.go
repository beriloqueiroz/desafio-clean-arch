package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"time"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/beriloqueiroz/desafio-clean-arch/configs"
	"github.com/beriloqueiroz/desafio-clean-arch/internal/event/handler"
	"github.com/beriloqueiroz/desafio-clean-arch/internal/infra/graph"
	"github.com/beriloqueiroz/desafio-clean-arch/internal/infra/grpc/pb"
	"github.com/beriloqueiroz/desafio-clean-arch/internal/infra/grpc/service"
	"github.com/beriloqueiroz/desafio-clean-arch/internal/infra/web/webserver"
	"github.com/beriloqueiroz/desafio-clean-arch/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig([]string{"."})
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Waiting rabbitmq...")
	time.Sleep(time.Second * 6)
	rabbitMQChannel := getRabbitMQChannel(configs.RabbitMqUrlConn)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})
	eventDispatcher.Register("OrderListed", &handler.OrderListedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrderUseCase := NewListOrderUseCase(db, eventDispatcher)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webListOrderHandler := NewWebListOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order", webOrderHandler.Create, "POST")
	webserver.AddHandler("/orders", webListOrderHandler.List, "GET")
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	orderService := service.NewOrderService(*createOrderUseCase, *listOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrderUseCase:   *listOrderUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

func getRabbitMQChannel(urlConn string) *amqp.Channel {
	conn, err := amqp.Dial(urlConn)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	q_created, err := ch.QueueDeclare(
		"orders_created", // name
		true,             // durable
		false,            // auto delete
		false,            // exclusive
		false,            // no wait
		nil,              // args
	)
	if err != nil {
		panic(err)
	}
	err = ch.QueueBind(q_created.Name, "order_created", "amq.direct", false, nil)

	if err != nil {
		panic(err)
	}

	q_listed, err := ch.QueueDeclare(
		"orders_listed", // name
		true,            // durable
		false,           // auto delete
		false,           // exclusive
		false,           // no wait
		nil,             // args
	)
	if err != nil {
		panic(err)
	}

	err = ch.QueueBind(q_listed.Name, "order_listed", "amq.direct", false, nil)

	if err != nil {
		panic(err)
	}

	return ch
}
