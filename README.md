# email-service

Service used for sending emails
Send trigger is incoming events from NATS
To be used for user registration, backup and restore, account management, etc.

## Test and Deploy

Use the built-in continuous integration in GitLab.

<h3>Local run</h3>
1. Using makefile

```make docker-compose-run```

2. Using docker compose directly

```	
docker compose -f deployment/docker/docker-compose.yml rm
docker compose -f deployment/docker/docker-compose.yml --env-file=.env --env-file=.env.credentials up --build --detach
```

***

## Installation
TBD

## Usage
TBD

## Support
TBD

## Contributing
TBD

## Authors and acknowledgment
TBD

## License
TBD

## Project status
TBD
