FROM golang:latest AS builder
LABEL maintainer="Nuttipong T<nuttipong.taechasanguanwong@allianz.com>"
WORKDIR /www
COPY go.mod go.sum ./
RUN go mod download
COPY . .
#RUN ls -alR
RUN go test -v -cover ./controllers && \
go test -v -cover ./services && \
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /www/main .
ARG WORK_ENV=development
ARG FROM_ENV="/www/config/config.${WORK_ENV}.json"
RUN echo ${FROM_ENV}
COPY --from=builder ${FROM_ENV} /root/config/
#RUN ls -alR
EXPOSE 8080
CMD ["./main"]
