# WarmStorageChain

WarmStorageChain is an open-source project that leverages blockchain technology to implement a heating Layer 2 solution for the Filecoin storage network. This project aims to enhance data reliability and efficiency by ensuring that stored data remains readily accessible and performs optimally.

## Introduction

### What is WarmStorageChain?

WarmStorageChain is designed to address the common challenges faced by decentralized storage systems, such as latency and data accessibility. By introducing a heating Layer 2 solution, WarmStorageChain ensures that data stored on the Filecoin network is kept in a 'warm' state, meaning it is quickly retrievable and consistently performant. The project integrates Cosmos SDK, Filecoin, and the Application Blockchain Interface (ABCI) to create a robust and efficient decentralized storage solution.

### Key Components

- **Cosmos SDK**: A framework for building custom, interoperable blockchain applications on top of Tendermint BFT.
- **Filecoin**: A decentralized storage network that turns cloud storage into an algorithmic market.
- **Heating Layer 2**: A secondary layer designed to keep frequently accessed data readily available, reducing latency and improving access times.
- **ABCI Layer**: Application Blockchain Interface layer to connect the Cosmos blockchain with external processes such as Filecoin retrieval tools.
- **RAFT Consensus Algorithm**: Ensures high availability and data consistency across private network nodes.

## Features

- **Heating Layer2 chain**: Keeps Filecoin-stored data in an optimal state for quick access and reliability, addressing the cold storage problem.
- **Filecoin Integration**: Leverages Filecoin for decentralized, secure, and scalable file storage.
- **Enhanced Efficiency**: Reduces latency and improves data retrieval times through an innovative heating mechanism.
- **Interoperability**: Ensures seamless interaction between different blockchain networks within the Cosmos ecosystem.
- **High Availability**: Uses RAFT consensus algorithm in a private network to ensure data availability and consistency.

## How It Works

### The Heating Layer 2 Concept

In a decentralized storage system like Filecoin, data is typically stored in a manner that prioritizes security and redundancy. However, this can sometimes result in slower access times, especially for data that is not frequently accessed. The heating Layer 2 in WarmStorageChain mitigates this by maintaining a subset of the data in a 'warm' state, ensuring that it can be accessed quickly when needed.

### Integration with Cosmos, Filecoin, and ABCI

- **Cosmos SDK**: The Cosmos SDK provides the tools to build a custom blockchain that can interact with other blockchains in the Cosmos ecosystem. This interoperability is crucial for ensuring that WarmStorageChain can interact seamlessly with other decentralized applications and services.
- **Filecoin**: Filecoin is used for its robust, decentralized file storage capabilities. The heating Layer 2 adds an additional layer of optimization on top of Filecoin's storage, ensuring that frequently accessed data is kept warm and quickly accessible.
- **ABCI Layer**: The Application Blockchain Interface (ABCI) layer connects the Cosmos blockchain with external processes, enabling integration with Filecoin retrieval tools. This allows the system to fetch Content Identifiers (CIDs) from Filecoin and store the retrieved data on the blockchain nodes.
- **RAFT Consensus Algorithm**: RAFT ensures that data is consistently replicated across nodes in the private network, providing high availability and reliability.

### Workflow

1. **Data Storage**: Data is stored on the Filecoin network using its decentralized storage protocol.
3. **Data Retrieval via ABCI**: When data needs to be accessed, the ABCI layer retrieves the data using Filecoin's retrieval tools based on the Content Identifier (CID).
4. **Heating Process**: Frequently accessed data is moved to the heating layer, ensuring it remains in a warm state for quick retrieval.
5. **Blockchain Storage**: The retrieved data is stored in the memory and disk storage of the blockchain nodes, ensuring fast access. The data storage utilizes the RAFT consensus algorithm to maintain consistency and high availability.
6. **Data Retrieval**: When a data request is made, the heating Layer2 ensures that the data is quickly retrieved, reducing latency and improving efficiency.

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. 

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

## Contributing

We welcome contributions from the community. To contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
