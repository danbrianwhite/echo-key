FROM alpine:3.3
ARG KEY

ENV PORT 8080
EXPOSE 8080

COPY main.go ./

RUN \
apk --no-cache add go && \
CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo --ldflags="-s -w" -o echo-key . && \
apk del go

CMD ["sh", "-c", "/echo-key -key ${key}"]
