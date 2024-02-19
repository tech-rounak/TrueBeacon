FROM golang:1.19-alpine AS go-build

WORKDIR /build/src/server
RUN apk add --no-cache gcc musl-dev
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ .
RUN CGO_ENABLED=0 GOOS=linux go build
ENV PORT=8000
ENV SECRET=true-beacon
CMD ["go", "run", "main.go"]
