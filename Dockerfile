FROM  golang:1.21.5-alpine3.19 as build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o /handle ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go

FROM scratch

COPY --from=build /handle /handle
COPY --from=build /app/.env /

ENTRYPOINT ["/handle"]