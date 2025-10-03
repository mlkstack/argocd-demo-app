FROM golang:1-alpine AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY go.mod ./
RUN go mod download

COPY main.go ./
RUN go build -ldflags '-w -s' -a -o /app main.go


FROM scratch
COPY --from=builder --chown=555 /app /app
USER 65532:65532
CMD [ "/app" ]