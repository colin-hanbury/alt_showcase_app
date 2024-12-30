FROM golang:1.19
WORKDIR /app
COPY go.mod go.sum ./
RUN npm i 
COPY . .
RUN go mod download
EXPOSE 8080
CMD ["/docker-gs-ping"]