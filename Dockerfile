FROM golang:1.18 as builder
WORKDIR /src
COPY . .
RUN git config --global url."https://theeraprasert:NHcfDjEynj4yAUsNDEKC@bitbucket.org".insteadOf "https://bitbucket.org"
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -o go ./

FROM alpine:3.14
RUN apk update && apk add bash
WORKDIR /
COPY --from=builder /src/go app

ENV TZ=Asia/Bangkok
CMD [ "./app" ]