FROM golang:1.23.12-alpine3.22

ENV SERVER_PORT="8081"
ENV ENV="dev"

COPY ./grpc-order-server .

RUN chmod +x grpc-order-server

EXPOSE 8081

ENTRYPOINT [ "sleep", "1h" ]
