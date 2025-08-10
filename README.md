# DID API MVP

A RESTful API for Decentralized Identifiers (DIDs) built with Go, implementing W3C DID specifications with Ed25519 cryptographic key pairs.

## 🚀 Features

- **DID Creation**: Generate new Decentralized Identifiers with Ed25519 key pairs
- **DID Resolution**: Retrieve DID documents by ID
- **Key Management**: Add verification keys to existing DIDs
- **Digital Signatures**: Sign and verify messages using Ed25519 cryptography
- **W3C Compliant**: Follows W3C DID specification standards
- **In-Memory Storage**: Fast, lightweight storage for development and testing

## Why We Use DIDs

| Reason | Explanation |
|--------|-------------|
| **Self-Sovereign Identity (SSI)** | DIDs allow individuals or entities to own and control their digital identity without a central authority. |
| **Privacy & Security** | No need to share sensitive data with intermediaries — you prove identity via cryptographic signatures. |
| **Interoperability** | DIDs follow the **W3C specification**, meaning they can work across different platforms, blockchains, and apps. |
| **Trust Without Centralization** | Blockchain or distributed systems ensure tamper-proof verification without a single point of failure. |
| **Portability** | You can carry your DID anywhere; it’s not tied to one company or platform. |

---

## Applications of DIDs

Your **DID API MVP** can be used in:

### 🔹 Authentication Systems
Replace usernames/passwords with DID-based cryptographic login.

### 🔹 Digital Credential Issuance
Universities, governments, or companies issue verifiable credentials tied to DIDs.

### 🔹 Secure Messaging
Sign messages so recipients can verify authenticity without intermediaries.

### 🔹 IoT Device Identity
Each IoT device gets a DID to prove authenticity and prevent spoofing.

### 🔹 Healthcare
Patients control medical records and share them securely with providers.

---

## 📁 Project Structure

```
did-api-mvp/
├── cmd/                    # CLI commands
│   ├── root.go            # Root command configuration
│   └── serve.go           # Server start command
├── config/                # Configuration management
│   └── config.go          # App, DB, and Logger configuration
├── conn/                  # Database connections
│   └── db.go              # Database connection setup
├── contracts/             # Smart contracts
│   └── did_contract.go    # DID contract definitions
├── controllers/           # HTTP request handlers
│   └── did_controller.go  # DID API endpoints
├── domain/                # Domain models
│   ├── did.go             # DID document structures
│   └── errors.go          # Custom error types
├── routes/                # Route definitions
│   └── routes.go          # API route setup
├── server/                # Server configuration
│   └── server.go          # HTTP server setup
├── services/              # Business logic
│   ├── crypto/            # Cryptographic operations
│   │   └── keypair.go     # Key generation and signing
│   └── did/               # DID operations
│       └── did_service.go # DID business logic
├── smartcontract/         # Blockchain contracts
│   └── DIDRegistry.sol    # Solidity DID registry
├── utils/                 # Utility functions
│   ├── errorutil/         # Error handling utilities
│   └── msgutil/           # Message utilities
├── go.mod                 # Go module definition
├── go.sum                 # Dependency checksums
├── main.go                # Application entry point
├── setup-env.ps1          # Environment setup script
└── load-env.ps1           # Environment loading script
```

## 🛠️ Prerequisites

- **Go 1.24.2 or higher**
- **Git** for version control
- **PowerShell** (for Windows environment scripts)

## ⚙️ Installation

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd did-api-mvp
   ```

2. **Install Go dependencies:**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables:**
   ```powershell
   # On Windows PowerShell
   powershell.exe -ExecutionPolicy Bypass -File .\setup-env.ps1
   
   # Or set manually:
   $env:APP_NAME = "did-api"
   $env:APP_PORT = "8080"
   $env:LOG_LEVEL = "debug"
   $env:LOG_FILE_PATH = "logs/app.log"
   ```

## 🚦 Running the Application

### Start the API Server

```bash
# Method 1: Using Go run
go run main.go serve

# Method 2: Build and run
go build -o did-api
./did-api serve
```

The server will start on `http://localhost:8080`

### Verify Server is Running

```bash
curl http://localhost:8080/ping
# Expected response: pong
```

## 📚 API Endpoints

### Base URL: `http://localhost:8080`

### 1. Health Check
```http
GET /ping
```
**Response:**
```
pong
```

---

### 2. Create DID
```http
POST /v1/api/did/create
```

**Response:**
```json
{
  "did": "did:key:5D2mPP4w8zGxJGGFKrP8c7Y1vT3zBnKxQmRtX8sC9dR1A2s3",
  "publicKey": "5D2mPP4w8zGxJGGFKrP8c7Y1vT3zBnKxQmRtX8sC9dR1A2s3",
  "privateKey": "3YB8jKRNG2PqC8nzMzRvQ1pX7tS4wY6nF9kL5dA8bV2mE1xR",
  "document": {
    "@context": "https://www.w3.org/ns/did/v1",
    "id": "did:key:5D2mPP4w8zGxJGGFKrP8c7Y1vT3zBnKxQmRtX8sC9dR1A2s3",
    "verificationMethod": [
      {
        "id": "did:key:5D2mPP4w8zGxJGGFKrP8c7Y1vT3zBnKxQmRtX8sC9dR1A2s3#keys-1",
        "type": "Ed25519VerificationKey2020",
        "controller": "did:key:5D2mPP4w8zGxJGGFKrP8c7Y1vT3zBnKxQmRtX8sC9dR1A2s3",
        "publicKeyBase58": "5D2mPP4w8zGxJGGFKrP8c7Y1vT3zBnKxQmRtX8sC9dR1A2s3"
      }
    ],
    "authentication": [
      "did:key:5D2mPP4w8zGxJGGFKrP8c7Y1vT3zBnKxQmRtX8sC9dR1A2s3#keys-1"
    ]
  }
}
```

---

### 3. Resolve DID
```http
GET /v1/api/did/resolve/{did-id}
```

**Example:**
```bash
curl http://localhost:8080/v1/api/did/resolve/did:key:5D2mPP4w8zGxJGGFKrP8c7Y1vT3zBnKxQmRtX8sC9dR1A2s3
```

**Response:**
```json
{
  "@context": "https://www.w3.org/ns/did/v1",
  "id": "did:key:5D2mPP4w8zGxJGGFKrP8c7Y1vT3zBnKxQmRtX8sC9dR1A2s3",
  "verificationMethod": [...],
  "authentication": [...]
}
```

---

### 4. Add Key to DID
```http
POST /v1/api/did/add-key/{did-id}
Content-Type: application/json
```

**Request Body:**
```json
{
  "publicKey": "3YB8jKRNG2PqC8nzMzRvQ1pX7tS4wY6nF9kL5dA8bV2m",
  "type": "Ed25519VerificationKey2020",
  "keyId": "keys-2"
}
```

**Response:**
Updated DID document with the new verification method.

---

### 5. Sign Message
```http
POST /v1/api/did/sign
Content-Type: application/x-www-form-urlencoded
```

**Form Data:**
- `message`: The message to sign
- `privateKey`: Base58-encoded private key

**Example:**
```bash
curl -X POST http://localhost:8080/v1/api/did/sign \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "message=Hello World" \
  -d "privateKey=3YB8jKRNG2PqC8nzMzRvQ1pX7tS4wY6nF9kL5dA8bV2mE1xR"
```

**Response:**
```json
{
  "message": "Hello World",
  "signature": "4Kx9mPp2wZ8rGbY5tN3vQ7sL1dF6hR9cA2eX8nT4yU6mV3k"
}
```

---

### 6. Verify Message
```http
POST /v1/api/did/verify
Content-Type: application/x-www-form-urlencoded
```

**Form Data:**
- `message`: The original message
- `signature`: Base58-encoded signature
- `publicKey`: Base58-encoded public key

**Example:**
```bash
curl -X POST http://localhost:8080/v1/api/did/verify \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "message=Hello World" \
  -d "signature=4Kx9mPp2wZ8rGbY5tN3vQ7sL1dF6hR9cA2eX8nT4yU6mV3k" \
  -d "publicKey=5D2mPP4w8zGxJGGFKrP8c7Y1vT3zBnKxQmRtX8sC9dR1A2s3"
```

**Response:**
```json
{
  "message": "Hello World",
  "signature": "4Kx9mPp2wZ8rGbY5tN3vQ7sL1dF6hR9cA2eX8nT4yU6mV3k",
  "valid": true
}
```

## 🧪 Testing

### Manual Testing with cURL

#### 1. Test DID Creation and Resolution
```bash
# Step 1: Create a new DID
curl -X POST http://localhost:8080/v1/api/did/create

# Step 2: Copy the DID from response and resolve it
curl http://localhost:8080/v1/api/did/resolve/did:key:{your-did-id}
```

#### 2. Test Message Signing and Verification
```bash
# Step 1: Create a DID and note the privateKey and publicKey
CREATE_RESPONSE=$(curl -s -X POST http://localhost:8080/v1/api/did/create)
PRIVATE_KEY=$(echo $CREATE_RESPONSE | jq -r '.privateKey')
PUBLIC_KEY=$(echo $CREATE_RESPONSE | jq -r '.publicKey')

# Step 2: Sign a message
SIGN_RESPONSE=$(curl -s -X POST http://localhost:8080/v1/api/did/sign \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "message=Hello DID World" \
  -d "privateKey=$PRIVATE_KEY")

SIGNATURE=$(echo $SIGN_RESPONSE | jq -r '.signature')

# Step 3: Verify the signature
curl -X POST http://localhost:8080/v1/api/did/verify \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "message=Hello DID World" \
  -d "signature=$SIGNATURE" \
  -d "publicKey=$PUBLIC_KEY"
```

### Testing with Postman

1. **Import Collection**: Create a new Postman collection with the above endpoints
2. **Environment Variables**: Set `{{base_url}}` to `http://localhost:8080`
3. **Test Workflow**:
   - Create DID → Save `did`, `publicKey`, `privateKey` as variables
   - Resolve DID using saved `did`
   - Sign message using saved `privateKey`
   - Verify signature using saved `publicKey` and signature

### Automated Testing

Create test files in Go:

```bash
# Create a test file
touch controllers/did_controller_test.go
touch services/did/did_service_test.go
```

Run tests:
```bash
go test ./...
```

## 🐛 Troubleshooting

### Common Issues

1. **Port already in use**:
   ```bash
   # Change the port in environment variables
   $env:APP_PORT = "8081"
   ```

2. **Go not found**:
   - Install Go from https://golang.org/dl/
   - Add Go to your PATH

3. **Permission denied on scripts**:
   ```powershell
   Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
   ```

### Logs

Check application logs:
```bash
# Logs are written to the path specified in LOG_FILE_PATH
cat logs/app.log
```

## 🔧 Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `APP_NAME` | `did-api` | Application name |
| `APP_PORT` | `8080` | Server port |
| `LOG_LEVEL` | `debug` | Logging level |
| `LOG_FILE_PATH` | `logs/app.log` | Log file location |
| `DB_HOST` | `localhost` | Database host |
| `DB_PORT` | `5432` | Database port |
| `DB_USER` | `postgres` | Database username |
| `DB_PASS` | `password` | Database password |
| `DB_SCHEMA` | `did_db` | Database schema |

## 🏗️ Architecture

### Key Components

1. **Controllers**: Handle HTTP requests and responses
2. **Services**: Implement business logic
3. **Domain**: Define data structures and models
4. **Config**: Manage application configuration
5. **Routes**: Define API endpoints and middleware

### Data Flow

```
HTTP Request → Routes → Controllers → Services → Domain Models → Response
```

### Cryptography

- **Algorithm**: Ed25519 (EdDSA)
- **Encoding**: Base58 for keys and signatures
- **Key Size**: 32 bytes (256 bits)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🔗 References

- [W3C DID Specification](https://www.w3.org/TR/did-core/)
- [DID Method: key](https://w3c-ccg.github.io/did-method-key/)
- [Ed25519 Signature Algorithm](https://tools.ietf.org/html/rfc8032)
- [Echo Web Framework](https://echo.labstack.com/)
- [Cobra CLI Library](https://cobra.dev/)

---

**Happy coding! 🚀**
