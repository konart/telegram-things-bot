FROM golang:1.9 AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/konart/telegram-things-bot
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

# We need certs to do secure requests
FROM centurylink/ca-certs
COPY --from=builder /app ./

# Cert and a key for webhhok
COPY server.crt /
COPY server.key /

ENTRYPOINT ["./app"]
ENV TELETHINGS_BOT_TOKEN=535194819:AAHg5VYEspALeV6bjKXF14nQbN8kSaPlAyA
ENV PORT=8080
CMD app