# Clean Architecture in Go

![ca.png](ca.png)

This repository provides an example implementation of Clean Architecture in Go. The Clean Architecture, popularized by Robert C. Martin (Uncle Bob), emphasizes separation of concerns, allowing for highly maintainable and testable code. This architecture enforces a structure that isolates different parts of the application, such as the domain, use cases, and external systems.

```azure
clean-arch/
├── cmd/
│   └── app/
│       └── main.go          
├── internal/
│   ├── app/
│   │   ├── client/
│   │   │   └── client.go    // HTTP client
│   │   ├── usecase/
│   │   │   └── room.go       // Use case 
│   │   └── repository/
│   │       └── room.go       // Repository
│   ├── domain/
│   │   └── room.go           // Entity
├── pkg/
│   └── config/
│       └── config.go         // Configuration management
├── api/
│   └── v1/
│       └── room.go           // API
├── go.mod                    
└── go.sum                   

```

## Key Features
- **Separation of Concerns:** Clearly separates different aspects of the application, such as business logic, data access, and presentation.
- **Testability:** Each component can be tested in isolation, ensuring the reliability and robustness of the application.
- **Flexibility:** Easily replace or update individual components without affecting the rest of the system.

## Contributing
Contributions are welcome! Feel free to submit a pull request or open an issue to discuss improvements, bug fixes, or new features.

## Running the Application

To run the application, follow these steps:

```shell
docker-compose -f docker-compose.yml up -d
go run cmd/app/main.go
```
