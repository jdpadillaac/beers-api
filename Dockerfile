FROM golang:1.17
WORKDIR /api/beers
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/beers /api/beers/cmd
EXPOSE ${PORT}
CMD ["./out/beers"]