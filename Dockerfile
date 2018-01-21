FROM golang:latest

RUN mkdir /app

ADD Scheduler_* /app/

WORKDIR /app

CMD ["/app/Scheduler_linux_amd64"]