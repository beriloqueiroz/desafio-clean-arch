FROM  golang:1.21.5-alpine3.19 as build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o /handle ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN ln -s /go/bin/linux_amd64/migrate /usr/local/bin/migrate

FROM scratch

COPY --from=build /handle /handle

ENTRYPOINT ["/handle"]