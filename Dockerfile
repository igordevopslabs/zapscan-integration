FROM golang:1.21 as builder

#avoid root
ENV USER=appuser 
ENV UID=1000

#avoid root
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy && CGO_ENABLED=0 GO111MODULES=on go build -ldflags="-s -w" -o zap-integration main.go

FROM scratch

#avoid root
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

WORKDIR /app
COPY --from=builder /app/zap-integration .
#avoid root
USER appuser:appuser

EXPOSE 9000

CMD ["./zap-integration"]