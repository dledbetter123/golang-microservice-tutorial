# Traditional Go Project Structure

```plaintext
docker --version
```

```plaintext
/cloud-native-app
├── cmd
│   ├── serviceA
│   │   └── main.go    # Entry point for the Service A microservice
│   └── serviceB
│       └── main.go    # Entry point for the Service B microservice
├── internal
│   ├── serviceA       # Internal code for service A
│   │   ├── handler
│   │   ├── service
│   │   └── repository
│   └── serviceB       # Internal code for service B
│       ├── handler
│       ├── service
│       └── repository
├── pkg
│   ├── common         # Common libraries used across your microservices
│   └── models         # Data models shared between services
├── web
│   ├── src            # React application source
│   ├── public
│   └── package.json   # Node.js specific config
├── configs            # Configuration file templates or default configs
├── scripts            # Scripts to perform various build, install, analysis, etc operations
├── deployments        # IaC for service deployments
├── Dockerfile
└── README.md
```