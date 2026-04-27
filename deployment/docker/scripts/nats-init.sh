#!/bin/sh

set -e

echo "🚀 Waiting for NATS..."

# wait until nats is ready
until nc -z nats 4222; do
  sleep 1
done

echo "📦 Creating JetStream stream: email_events"

nats --server nats://nats:4222 stream add email_events \
  --subjects "email.send,email.receive" \
  --storage file \
  --replicas 1 \
  --max-msgs 10000 \
  --max-age 24h || true

echo "✅ NATS initialized"