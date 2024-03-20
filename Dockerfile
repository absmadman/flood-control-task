FROM golang:alpine

ADD go.mod .

COPY . .

CMD ["go", "run", "main.go"]