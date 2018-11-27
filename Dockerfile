FROM golang:1.11.2-alpine AS build
RUN apk update && apk add git ca-certificates
WORKDIR /build
ENV GO111MODULE on
COPY go.mod .
COPY go.sum .
COPY . .
RUN CGO_ENABLED=0 go build -o graphql-go-workshop

FROM alpine:3.8
RUN apk update && apk add ca-certificates
WORKDIR /usr/bin
COPY --from=build /build/graphql-go-workshop .
EXPOSE 8080
CMD ["graphql-go-workshop"]