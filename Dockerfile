FROM golang:1.22-alpine
WORKDIR /app
COPY . /app
RUN go mod download && go mod verify
RUN go build -o app
CMD [ "./app" ]