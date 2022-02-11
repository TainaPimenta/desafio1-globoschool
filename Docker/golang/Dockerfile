#versão do golang que está sendo usada para o docker
FROM golang:1.17

#Utilizando a pasta onde se encontra o arquivo
RUN mkdir -p /app
WORKDIR /app
COPY . /app

#Para conseguir fazer com que o codigo rode dentro do docker
RUN go mod tidy