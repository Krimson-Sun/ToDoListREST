FROM golang:1.23

ENV GOPATH=/
ENV GOSUMDB="off"

COPY ./ ./

RUN go mod download

RUN go build -o sosiska .

CMD ["./sosiska"]