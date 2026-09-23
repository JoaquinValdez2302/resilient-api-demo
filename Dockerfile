# Etapa 1: Construcción
FROM golang:alpine AS builder
WORKDIR /app
# Copiamos los archivos de nuestro proyecto
COPY go.mod ./
COPY main.go ./
# Compilamos el binario de Go (lo llamaremos 'api')
RUN go build -o api main.go

# Etapa 2: Imagen final ultra ligera
FROM alpine:latest
WORKDIR /root/
# Copiamos SOLAMENTE el binario compilado de la etapa anterior
COPY --from=builder /app/api .
EXPOSE 8080
# Comando para ejecutar la aplicación
CMD ["./api"]