FROM golang:1.11.5

ENV APP_NAME desafio
ENV PORT 80

COPY . /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

RUN go get ./
RUN go build -o ${APP_NAME}

CMD ./${APP_NAME}

EXPOSE ${PORT}