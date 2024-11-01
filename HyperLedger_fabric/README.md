# Hyperledger Fabric Project

This guide will help you set up and run your Hyperledger Fabric network, deploy the asset transfer chaincode, and use a Node.js application to interact with the blockchain via RESTful APIs.

---

## Steps to Install and Setup the Network

### 1. Navigate to Your Project Directory

```bash
cd Hyperfiber-intial-edit
```

### 2. Install Hyperledger Fabric

```bash
curl -sSL https://bit.ly/2ysbOFE | bash -s -- 2.5.0
```

### 3. Go to the Fabric Test Network Directory

```bash
cd fabric-samples/test-network
```

### 4. Stop Any Existing Network

```bash
./network.sh down
```

### 5. Start the Network with CA Enabled and Create a Channel

```bash
sudo ./network.sh up createChannel -ca -c mychannel
```

### 6. Deploy the Chaincode

```bash
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-javascript -ccl javascript
```

### 7. Set Environment Variables for the Peer

```bash
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
```

### 8. Invoke the Chaincode

```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" \
  -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" \
  --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" \
  -c '{"function":"InitLedger","Args":[]}'
```

---

## Running the Node.js Application

### 1. Go to the Node.js Application Directory

```bash
cd asset-transfer-basic/application-gateway-javascript
```

### 2. Install Required Dependencies

```bash
npm i
```

### 3. Start the Node.js Server

```bash
npm start
```

The server should start at `http://localhost:3000`.

---

## Accessing and Testing the APIs

Once the server is running locally, you can access the API endpoints using the following examples:

### 1. Initialize Ledger

**Endpoint**: `POST /ledger/init`

**Example**:
```bash
POST http://localhost:3000/ledger/init
```

**Response**:
```json
{
  "message": "Ledger initialized successfully"
}
```

---

### 2. Get All Assets

**Endpoint**: `GET /assets`

**Example**:
```bash
GET http://localhost:3000/assets
```

**Response**:  
A list of assets in the ledger.

---

### 3. Create Asset (Example for Dealer 3)

**Endpoint**: `POST /asset`

**Example**:
```json
{
  "id": "asset3",
  "dealerId": "dealer3",
  "msisdn": "1122334455",
  "mpin": "91011",
  "balance": 500,
  "status": "active",
  "transAmount": 0,
  "transType": "",
  "remarks": ""
}
```

**Response**:
```json
{
  "message": "Asset asset3 created successfully"
}
```

---

### 4. Update Asset

**Endpoint**: `PUT /asset`

**Example**:
```json
{
  "id": "asset1",
  "dealerId": "dealer1",
  "msisdn": "1234567890",
  "mpin": "1234",
  "balance": 350,
  "status": "active",
  "transAmount": 50,
  "transType": "credit",
  "remarks": "Updated balance"
}
```

**Response**:
```json
{
  "message": "Asset asset1 updated successfully"
}
```

---

### 5. Transfer Asset

**Endpoint**: `POST /asset/transfer`

**Example**:
```json
{
  "id": "asset2",
  "newOwner": "dealer3"
}
```

**Response**:
```json
{
  "message": "Successfully transferred asset asset2 from dealer2 to dealer3"
}
```

---

### 6. Read Asset

**Endpoint**: `GET /asset/:id`

**Example**:
```bash
GET http://localhost:3000/asset/asset1
```

---

### 7. Get Asset Transaction History

**Endpoint**: `GET /asset/:id/history`

**Example**:
```bash
GET http://localhost:3000/asset/asset1/history
```

---

This structured guide provides steps to set up the Hyperledger Fabric network, deploy the asset transfer chaincode, and interact with the blockchain using RESTful APIs through the Node.js application.
