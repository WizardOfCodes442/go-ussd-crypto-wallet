module github.com/WizardOfCodes442/go-ussd-crypto-wallet

go 1.23

require (
	
	// REST
    github.com/gin-gonic/gin v1.9.0         // HTTP framework for REST APIs

    // GraphQL
    github.com/99designs/gqlgen v0.17.35    // GraphQL server library
    github.com/graphql-go/graphql v0.10.0  // Alternative GraphQL library

    // gRPC
    google.golang.org/grpc v1.58.0          // gRPC library
    google.golang.org/protobuf v1.31.0      // Protocol Buffers

    // Utility libraries
    github.com/joho/godotenv v1.5.1         // Environment variables
    github.com/stretchr/testify v1.8.1     // Testing utilities
    github.com/sirupsen/logrus v1.9.0      // Structured logging
	// Bitcoin and related libraries
	github.com/btcsuite/btcd v0.24.2
	github.com/btcsuite/btcutil v1.1.6
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0
	github.com/tyler-smith/go-bip39 v1.1.0

	// Ethereum and EVM-compatible chains
	github.com/ethereum/go-ethereum v1.11.5
	github.com/ethereum/go-ethereum/rpc v1.11.5 // RPC client for Ethereum

	// Blockchain node providers SDKs
	github.com/alchemy-platform/alchemy-sdk-go v1.0.3
	github.com/infura/infura-go-sdk v0.2.5
	github.com/blockcypher/gobcy v0.9.0
	github.com/moralisweb3/moralis-go-sdk v1.0.0

	// Additional chains
	github.com/portto/solana-go-sdk v1.7.0       // Solana
	github.com/centrifuge/go-substrate-rpc-client v0.9.0 // Polkadot/Substrate
	github.com/cosmos/cosmos-sdk v0.46.6         // Cosmos SDK-based chains
	github.com/stellar/go v4.7.1+incompatible    // Stellar
)

require (
	// Indirect dependencies
	github.com/btcsuite/btcd/btcec/v2 v2.1.3 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.1.0 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/sys v0.0.0-20200814200057-3d37ad5750ed // indirect
)
