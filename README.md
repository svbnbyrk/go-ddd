# Go-Driven Domain Design (DDD) Project

This project showcases the implementation of **Domain-Driven Design (DDD)** principles using Go, focusing on both tactical and strategic elements. The development process is **OpenAPI 3.0 documentation-first**, ensuring that the API is well-defined and consistent before implementation.

---

## Table of Contents

1. [Project Goals](#project-goals)
2. [Directory Structure](#directory-structure)
3. [Features](#features)
4. [Installation](#installation)
   - [Prerequisites](#prerequisites)
   - [Steps](#steps)
5. [How It Works](#how-it-works)
6. [Contributing](#contributing)
7. [License](#license)

---

## Project Goals

1. **Apply DDD Tactical and Strategic Design**:
   - Develop self-managing entities, aggregates, and domain services.
   - Establish bounded contexts and strategic communication patterns.

2. **Documentation-First Approach**:
   - Use OpenAPI 3.0 specifications to design and document APIs before implementation.
   - Automate code generation and testing based on the OpenAPI specification.

3. **Flexible and Maintainable Architecture**:
   - Focus on separation of concerns for scalability and clarity.

---

## Directory Structure

The project follows a modular structure for clarity and maintainability:

```plaintext
.
├── api               # OpenAPI 3.0 specifications and generated code
├── internal
│   ├── domain        # Core domain logic (entities, value objects, aggregates)
│   ├── ports         # Interfaces for external systems (e.g., repository, services)
│   ├── adapters      # Implementation of ports (e.g., database, API clients)
│   └── app           # Application services (use cases and workflows)
└── pkg               # Shared utilities, helpers, or common libraries
```

---

## Features

- **Domain-Centric Design**:
  - Encapsulate business logic within the `domain` layer.
  - Use `ports` and `adapters` for clean dependency inversion.

- **OpenAPI 3.0 Documentation-First**:
  - Define APIs in the `api` directory using `api.yml`.
  - Generate server stubs and client code using tools like `oapi-codegen`.

- **Extensibility and Modularity**:
  - Decouple domain logic from infrastructure to facilitate changes.

---

## Installation

### Prerequisites

- Go 1.23+ (consider upgrading to 1.23)
- Tools for OpenAPI code generation (e.g., `oapi-codegen`)

### Steps

1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd <repository_name>
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Generate code from OpenAPI specs:
   ```bash
   oapi-codegen -generate types -o internal/app/api_types.gen.go api/api.yml
   oapi-codegen -generate chi-server -o internal/app/api_server.gen.go api/api.yml
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

---

## How It Works

1. **OpenAPI Specification**:
   - Define API contracts in `api/api.yml`.

2. **Code Generation**:
   - Use `oapi-codegen` to generate server interfaces, request/response types, and handlers.

3. **DDD Principles**:
   - Domain entities are located in `internal/domain`.
   - Communication between layers happens through `ports` and their implementations in `adapters`.

4. **Implementation**:
   - Develop application logic in the `app` layer, leveraging domain logic and adapters.

---

## Contributing

Contributions are welcome! Follow these steps:

1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature/<feature_name>
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add <feature_name>"
   ```
4. Push to your branch:
   ```bash
   git push origin feature/<feature_name>
   ```
5. Open a pull request.

---

## License

This project is licensed under the [MIT License](LICENSE).

---
