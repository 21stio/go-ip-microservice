FROM golang:latest

RUN curl https://glide.sh/get | sh

WORKDIR /go/src/github.com/21stio/go-ip-microservice
COPY . .

RUN glide install
RUN go build -o application

RUN ./downloadGeoipDatabase.sh

CMD ./application -APPLICATION_PORT=${APPLICATION_PORT} -CRON_SCHEDULE=${CRON_SCHEDULE}