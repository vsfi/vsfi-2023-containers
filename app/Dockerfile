FROM registry-01.vsfi.org/golang:1.20-alpine as build
WORKDIR /app
# copy files with depends
COPY go.mod go.sum ./
# install depends
RUN go mod download
# copy source code
COPY *.go ./

# compile app
RUN CGO_ENABLED=0 GOOS=linux go build -o beer

FROM registry-01.vsfi.org/alpine:3
WORKDIR /app

COPY --from=build /app/beer .

COPY templates/ templates/

ENTRYPOINT [ "/app/beer" ]