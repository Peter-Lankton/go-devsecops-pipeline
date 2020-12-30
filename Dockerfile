FROM golang:1.14-alpine
LABEL base.name="goDevSecOpsDocker"
WORKDIR /app
COPY . . 
RUN go build -o main .
RUN chmod +x main
EXPOSE 8080
ENTRYPOINT ["./main"]