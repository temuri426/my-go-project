FROM  golang:1.24-bookworm
WORKDIR /app
COPY api.go .
COPY go.mod .
EXPOSE 8090
CMD ["go","run","api.go"]
