# Use a imagem oficial do Golang
FROM golang:1.17

# Define o diretório de trabalho
WORKDIR /app

# Copia o código-fonte para o contêiner
COPY . .

# Define o comando de inicialização
CMD ["./api_rest_gin_go"]
