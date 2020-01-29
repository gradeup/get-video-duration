FROM golang:1.13-alpine as builder
WORKDIR /usr/src/app
RUN apk add --no-cache git ca-certificates tzdata && update-ca-certificates 

# Create an unpriveleged user 
RUN adduser -D -g '' gopher
# Install deps
COPY go.mod .
COPY go.sum .
RUN go mod download

# Add source
ADD . .

# Build the binary
RUN CGO_ENABLED=0  go build -o duration *.go

FROM alpine
# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
COPY --from=builder /usr/src/app/duration /duration
RUN apk add ffmpeg
# Use an unprivileged user.
USER gopher

CMD ["/duration"]
