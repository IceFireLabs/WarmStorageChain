# WarmStorageChain

WarmStorageChain is an innovative open-source project designed to address the common challenges faced by decentralized storage systems, specifically focusing on data latency and accessibility. By leveraging blockchain technology, WarmStorageChain implements a heating Layer2 solution for the Filecoin storage network. This layer ensures that data remains in a 'warm' state, meaning it is quickly retrievable and consistently performant.

## I. Overview

Decentralized storage systems like Filecoin offer robust, secure, and scalable storage solutions by distributing data across a network of nodes. However, the decentralized nature of these systems often leads to increased latency and slower access times, especially for data that is not frequently accessed (cold storage). This latency can significantly impact the performance and user experience of applications relying on decentralized storage.

## II. Solution: Hot Storage and Retrieval Tier

WarmStorageChain introduces a heating Layer2 solution to mitigate the latency issues inherent in decentralized storage systems. The heating Layer2 ensures that frequently accessed data remains in a 'warm' state, providing quicker retrieval times and improved overall performance. This solution is achieved through the integration of several key technologies:

- **Filecoin**: A decentralized storage network that transforms cloud storage into an algorithmic market. Filecoin's robust and secure storage capabilities serve as the foundation for WarmStorageChain's data storage.
- **Tendermint ABCI (Application Blockchain Interface)**: The ABCI layer connects the blockchain with external processes, such as Filecoin retrieval tools. This integration allows WarmStorageChain to fetch Content Identifiers (CIDs) from Filecoin and store the retrieved data on blockchain nodes.
- **Lassie**: Filecoin's file retrieval tool, which is used to efficiently retrieve data from the Filecoin network based on the provided CIDs.

## III. Key Features

- **Filecoin Integration**: Leverages Filecoin for decentralized, secure, and scalable file storage.
- **Heating Layer2 Chain**: Keeps frequently accessed data in an optimal state for quick access, addressing the cold storage problem and enhancing data reliability.
- **Enhanced Efficiency**: Reduces latency and improves data retrieval times through an innovative heating mechanism.

## IV. Detailed Workflow

1. **Data Storage**: Data is initially stored on the Filecoin network using its decentralized storage protocol.
2. **Data Retrieval via Tendermint ABCI and Filecoin Lassie Tool**: When a user submits a CID from the Filecoin storage network, WarmStorageChain receives this CID and begins the retrieval process.
3. **Payment Processing**: The user pays a fee for the retrieval service. Once the payment is confirmed, the blockchain ABCI function is triggered.
4. **Data Retrieval Using Lassie**: The ABCI layer uses Filecoin's Lassie retrieval tool to fetch the data associated with the CID.
5. **Storing Retrieved Data**: Upon successful retrieval, WarmStorageChain stores the data on the disk storage of the blockchain nodes.
6. **Transaction Logging**: The retrieval and storage transaction is recorded on the blockchain, ensuring transparency and traceability.
7. **Data Access**: When a data request is made, the heating Layer2 ensures that the data is quickly retrieved from the blockchain nodes' storage, reducing latency and improving efficiency.

## V. Integration and Interoperability

WarmStorageChain's integration with Filecoin allows it to leverage Filecoin's decentralized storage capabilities effectively. The use of the Tendermint ABCI layer ensures that WarmStorageChain can seamlessly interact with Filecoin retrieval tools, enabling efficient data management and access.

## VI. Benefits

- **Reduced Latency**: By maintaining frequently accessed data in a warm state, WarmStorageChain significantly reduces data retrieval times.
- **Enhanced Interoperability**: Uses the ABCI protocol to interact seamlessly with Filecoin and other blockchain networks.

## VII. How It's Made

### VII.I Technologies Used

- **Filecoin**: Primary storage layer, providing secure and scalable storage.
- **Tendermint ABCI**: Customized interaction with Filecoin and Lassie.
- **Lassie (Filecoin's Tool)**: Efficiently retrieves data based on CIDs.

### VII.II Architecture

The following diagram illustrates the WarmStorageChain architecture:

```
+---------------------+     +--------------------+
|  User CID Submission|     |  Payment Processing|
|      (Client)       |     |      (Client)      |
+----------+----------+     +----------+---------+
           |                           |
           |                           |
           v                           v
+----------+----------+     +----------+--------------+
| WarmStorageChain API|     |     ABCI Function       |
| (Filecoin CID & Fee)|     |(Trigger filecoin Lassie)|
+----------+----------+     +----------+--------------+
           |                           |
           v                           v
+----------+----------+     +----------+---------+
| Lassie Retrieval Tool|    |  Filecoin Storage  |
|  (Fetch Data via CID)|    |  (Decentralized)   |
+----------+----------+     +----------+---------+
           |                           |
           v                           v
+----------+----------+     +----------+---------+
|   Store Data on      |    | Transaction Logging|
| WarmStorageChain Node|    | (Blockchain Record)|
|    (Heating Layer2)  |    |                    |
+----------+----------+     +----------+---------+
           |
           v
+----------+----------+
|  Gateway Access     |
|  (Retrieve Data)    |
+---------------------+
```

### VII.III Challenges and Solutions

- **Seamless Integration**: Customized Tendermint ABCI logic for efficient Filecoin-ABCI communication.
- **Data Storage Enhancement**: 'Hot Storage Tier' design improves access speed by storing data on blockchain nodes' disks.

### VII.IV Partner Technologies

- **Filecoin**: Robust decentralized storage foundation.
- **Tendermint ABCI**: Facilitates efficient external process integration.
- **Lassie**: Faster data retrieval, ensuring 'hot' data availability.

### VII.V Notable Innovations

- **Custom ABCI Logic**: Enhances interaction efficiency for rapid data retrieval.
- **'Hot Storage and Retrieval Tier' Design**: Optimizes Filecoin access speed by storing data internally on blockchain nodes' disks.

## VIII. Get Started

### VIII.I Prerequisites

Before you begin, ensure you have met the following requirements:
- You have installed [Golang](https://golang.org/doc/install).
- You have installed [Filecoin lassie](https://github.com/filecoin-project/lassie/)
- You have installed [cosmos ignite](https://ignite.com/).
- You have installed [node](https://nodejs.org/en/download/package-manager).
- You have installed [pnpm](https://pnpm.io/installation).

### VIII.II Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/BlockCraftsman/WarmStorageChain.git
    cd WarmStorageChain
    ```

2. Build and run blockchain:
    ```bash
    cd WarmStorageChain
    ignite chain serve
    ```

3. run blockchain client vue app:
    ```bash
    cd WarmStorageChain/vue
    pnpm install
    pnpm run dev
    ```

## IX. Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

### IX.I How to Contribute

1. Fork the Project.
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`).
3. Commit your Changes (`git commit -m -s 'Add some AmazingFeature'`).
4. Push to the Branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

### IX.II Code of Conduct

Please read the [Code of Conduct](CODE_OF_CONDUCT.md) to understand the expectations for contributors.

### IX.III License

Distributed under the MIT License. See `LICENSE` for more information.
