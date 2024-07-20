FROM golang:1.20-alpine AS builder

RUN adduser -D -g '' go_user

WORKDIR /src

COPY ./main.go .

RUN go build -o /app/mass_program main.go

FROM alpine:3.20 AS runner

RUN adduser -D -g '' go_user && \
    mkdir -p /app && \
    chown go_user:go_user /app && \
    chmod +x /app

COPY --from=builder /app/mass_program /app/mass_program

USER go_user
EXPOSE 8080

CMD ["/app/mass_program", "8080"]