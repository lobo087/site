# ===== Etapa 1: build =====
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Copiamos solo el código (este Dockerfile debe estar junto a main.go)
COPY . .

# Si no existe go.mod, lo creamos con un nombre simple
RUN [ -f go.mod ] || go mod init apigo

# Compilar el binario (tu main está en el directorio actual)
# CGO deshabilitado para binario estático y más portable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

# ===== Etapa 2: runtime mínimo =====
FROM alpine:3.20
WORKDIR /app

# Copiamos el binario desde la etapa de build
COPY --from=builder /app/server /app/server

# Puerto usado por tu API en Go (ajústalo si usas otro)
EXPOSE 8085

# (Opcional) Healthcheck básico
HEALTHCHECK --interval=30s --timeout=3s --retries=3 CMD wget -qO- http://127.0.0.1:8085/libros || exit 1

CMD ["/app/server"]
