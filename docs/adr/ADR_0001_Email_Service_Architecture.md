# Event-Driven Email Service Architecture (NATS-based)

> **Status**: Accepted  
> **Date**: 2026-04-27  
> **Related Tickets**: KIVP-XXX

---

#### Context

We need a dedicated service responsible for sending emails triggered by system events such as user registration, password resets, backups, and account management actions.

Direct service-to-service communication (e.g., HTTP calls) would introduce tight coupling, reduce resilience, and make scaling and failure handling more difficult.

To improve decoupling and reliability, we require an asynchronous event-driven approach where email delivery is handled independently from upstream services.

---

#### Decision

We implement a standalone **email-service** based on an event-driven architecture using **NATS** as the message broker.

Key design decisions:

- All email-triggering events are published to NATS.
- The email-service subscribes to the `email.receive` subject.
- The service processes incoming events and sends emails via configured SMTP/email provider.
- The service is stateless; all required data must be included in the event payload.

### Example event:

```json
{
  "type": "user.registration",
  "to": "user@example.com",
  "subject": "Welcome!",
  "template": "welcome-email",
  "data": {
    "username": "john_doe"
  }
}
```

### NATS subject:
```
email.receive
```

### Manual test (Windows PowerShell):
```
Get-Content .\email-event.json -Raw | nats pub email.receive
```
#### Consequences

| Positive | Negative |
|---------|----------|
| ✅ Loose coupling between services	| ❌ Requires proper event schema management |
| ✅ Asynchronous processing improves resilience	| ❌ Debugging requires tracing through message broker |
| ✅ Easy horizontal scaling	| ❌ Must handle duplicate events (idempotency required) |
| ✅ Independent deployment lifecycle	| ❌ Slight operational overhead (NATS dependency) |

#### Next Steps

- [ ] Define strict event schema contracts (versioned)
- [ ] Implement idempotency handling for duplicate messages
- [ ] Add structured logging for event tracing
- [ ] Introduce dead-letter / retry strategy for failed emails