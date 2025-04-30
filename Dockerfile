FROM golang:1.17 as builder
WORKDIR /go/src/github.com/udistrital/agora_api_crud
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates python3 py3-pip
RUN pip3 install awscli
WORKDIR /root/
COPY --from=builder /go/src/github.com/udistrital/agora_api_crud/main .
COPY entrypoint.sh .
COPY conf/app.conf conf/app.conf
RUN chmod +x main entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]