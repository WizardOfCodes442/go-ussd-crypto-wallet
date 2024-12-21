crypto_wallet_api/
├── cmd/
│   ├── server/
│   │   ├── main.go  # Entry point for the application
│   └── ussd/
│       ├── main.go  # USSD-specific entry point (if separate server required)
├── configs/
│   ├── app_config.yaml  # Application configuration file
│   └── secrets.env       # Environment secrets (gitignored)
├── internal/
│   ├── auth/
│   │   ├── handlers.go   # Handlers for authentication endpoints
│   │   ├── middleware.go # Authentication middleware
│   │   ├── models.go     # Models for user and authentication
│   │   └── services.go   # Authentication-related business logic
│   ├── wallet/
│   │   ├── handlers.go   # Wallet endpoints handlers
│   │   ├── models.go     # Models for wallet and related entities
│   │   ├── services.go   # Business logic for wallets
│   │   └── utils.go      # Wallet utility functions
│   ├── transactions/
│   │   ├── handlers.go   # Handlers for transaction-related endpoints
│   │   ├── models.go     # Models for transactions
│   │   └── services.go   # Transaction-related logic
│   ├── blockchain/
│   │   ├── clients.go    # Blockchain interaction clients (e.g., Ethereum, Bitcoin)
│   │   ├── services.go   # Blockchain-specific operations
│   └── notifications/
│       ├── handlers.go   # Handlers for notifications
│       ├── services.go   # Notification services
│       └── utils.go      # Notification utilities
├── pkg/
│   ├── utils/
│   │   ├── encryption.go # Utility functions for encryption
│   │   ├── validation.go # Utility functions for request validation
│   │   └── logger.go     # Logging utilities
│   ├── database/
│   │   ├── db.go         # Database initialization and connection
│   │   └── migrations/   # Database migration files
│   └── network/
│       ├── rest_client.go # HTTP client for external REST APIs
│       └── ussd_client.go # USSD network communication logic
├── api/
│   ├── docs/
│   │   └── swagger.yaml  # API documentation
│   ├── routes.go         # API routing setup
│   └── middleware.go     # General middleware (e.g., logging, CORS)
├── tests/
│   ├── auth_tests.go     # Unit tests for authentication
│   ├── wallet_tests.go   # Unit tests for wallets
│   ├── transaction_tests.go # Unit tests for transactions
│   └── integration/      # Integration tests
├── Makefile              # Commands for building, testing, running the app
├── go.mod                # Go module file
├── go.sum                # Dependencies checksum
└── README.md             # Project overview and instructions
