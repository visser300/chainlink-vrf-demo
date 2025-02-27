# Chainlink VRF Consumer Project


This project demonstrates how to integrate Chainlink's Verifiable Random Function (VRF) into your smart contracts. It provides a simple implementation of a random number consumer that can be deployed to various networks including local development environments, testnets like Sepolia, and mainnet.


## Project Structure

```
├── script/
│   └── DeployRandomNumberConsumer.s.sol  # Deployment script
├── src/
│   └── RandomNumberConsumer.sol          # Main contract
├── test/
│   └── SepoliaIntegrationTest.sol        # Integration test for Sepolia
├── lib/                                  # Dependencies
│   ├── chainlink/                        # Chainlink contracts
│   └── forge-std/                        # Forge standard library
├── cache/                                # Forge cache
└── out/                                  # Compiled output
```


## Tech Stack

Solidity: Smart contract programming language
Foundry: Development framework for Ethereum applications
    Forge: Testing framework
    Anvil: Local Ethereum node
    Cast: Ethereum transaction and RPC client
Chainlink VRF: Verifiable Random Function for secure on-chain randomness


### Smart Contracts

RandomNumberConsumer.sol: The main contract that consumes randomness from Chainlink VRF. It:
    Requests random numbers from the Chainlink VRF Coordinator
    Stores the received random values
    Provides functions to access the random values


## Configuration

Environment Variables
Create a .env file with the following variables:

```
PRIVATE_KEY=your_wallet_private_key
SUBSCRIPTION_ID=your_chainlink_vrf_subscription_id
KEYHASH=your_chainlink_vrf_keyhash
DEPLOYED_CONTRACT_ADDRESS=address_of_deployed_contract
DEPLOYER_ADDRESS=address_of_deployer
```

Network Configuration
The project supports multiple networks through Foundry's configuration. The StdChains.sol file includes RPC URLs for various networks including:
    Ethereum (Mainnet, Sepolia, Holesky)
    Layer 2s (Optimism, Arbitrum, Base)
    Sidechains (Polygon, Avalanche, BSC)
    And many others

Running Tests and Deployments
Local Testing
1. Start a local Anvil node

```
anvil --fork-url $RPC_URL --chain-id 11155111
```

Deploy the contract to your local node:

```
   source .env
   forge script script/DeployRandomNumberConsumer.s.sol --broadcast --rpc-url http://localhost:8545 --private-key $PRIVATE_KEY
```

Sepolia Testnet Deployment
Deploy to Sepolia:

```
   source .env
   forge script script/DeployRandomNumberConsumer.s.sol --broadcast --rpc-url https://eth-sepolia.g.alchemy.com/v2/YOUR_API_KEY --private-key $PRIVATE_KEY
```

Run integration tests against your deployed contract:

```
   forge test --match-contract SepoliaIntegrationTest --rpc-url https://eth-sepolia.g.alchemy.com/v2/YOUR_API_KEY -vvv
```

Mainnet Deployment
Deploy to Ethereum Mainnet (be careful with gas costs):
```
   source .env
   forge script script/DeployRandomNumberConsumer.s.sol --broadcast --rpc-url https://eth-mainnet.g.alchemy.com/v2/YOUR_API_KEY --private-key $PRIVATE_KEY
```

Chainlink VRF Configuration
Before using this contract, you need to:
1. Create a VRF subscription at vrf.chain.link
2. Fund your subscription with LINK tokens
3. Add your deployed contract address as a consumer to your subscription
4. Use the appropriate VRF Coordinator address and key hash for your network:
    Sepolia: Coordinator 0x8103B0A8A00be2DDC778e6e7eaa21791Cd364625
    For other networks, refer to the Chainlink documentation

Notes
1. The integration test in SepoliaIntegrationTest.sol is designed to work with a deployed contract on Sepolia
2. Make sure your wallet has enough ETH for gas fees and your subscription has enough LINK tokens
3. For production use, carefully review the randomness requirements and consider additional security measures