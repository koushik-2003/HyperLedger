# Hyperledger Fabric Project

This guide will help you set up and run your Hyperledger Fabric network, deploy the asset transfer chaincode, and use a Node.js application to interact with the blockchain via RESTful APIs.

---

## Steps to Install and Setup the Network

# **Hyperledger Fabric Project**

This guide will help you set up and run your Hyperledger Fabric network, deploy the asset transfer chaincode, and use a Node.js application to interact with the blockchain via RESTful APIs.

---

## **Steps to Install and Setup the Network**

### **1. Navigate to Your Project Directory**
cd HyperLedger_Fabric_Project/HyperLedger_fabric

### **2. Install Hyperledger Fabric**
curl -sSL https://bit.ly/2ysbOFE | bash -s -- 2.5.0

### **3. Go to the Fabric Test Network Directory**
cd fabric-samples/test-network

### **4. Grant Execute Permissions**: To make the script executable, use the following command:
chmod +x network.sh

### **5. Stop Any Existing Network**
./network.sh down

### **6. Start the Network with CA Enabled and Create a Channel**
sudo ./network.sh up createChannel -ca -c mychannel

### **7. Deploy the Chaincode**
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-javascript -ccl javascript

### **8. Set Environment Variables for the Peer**
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

### 9. Invoke the Chaincode

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
POST http://localhost:3000/ledger
```

**Response**:
```json
{
  "message": "Ledger initialized successfully"
}
```

---

### 2. Get All Assets

**Endpoint**: `GET /getassets`

**Example**:
```bash
GET http://localhost:3000/getassets
```

**Response**:  
A list of assets in the ledger.

---

### 3. Create Asset (Example for Dealer 3)

**Endpoint**: `POST /createasset`

**Example**:
```json
{
  "ID": "adluru",
  "DEALERID": "koushik",
  "MSISDN": "741",
  "MPIN": "2001",
  "BALANCE": 300,
  "STATUS": "active",
  "TRANSAMOUNT": 0,
  "TRANSTYPE": "",
  "REMARKS": ""
}
```

**Response**:
```json
{
  "message": "Asset adluru created successfully"
}
```

---

### 4. Update Asset

**Endpoint**: `PUT /updateasset`

**Example**:
```json
{
  "id": "shiva",
  "dealerId": "sale",
  "msisdn": "258",
  "mpin": "123",
  "balance": 3625,
  "status": "active",
  "transAmount": 60,
  "transType": "debit",
  "remarks": "update balance"
}
```

**Response**:
```json
{
  "message": "Asset shiva updated successfully"
}
```

---

### 5. Transfer Asset

**Endpoint**: `POST /asset/transfer`

**Example**:
```json
{
  "id": "shiva",
  "newOwner": "malli"
}
```

**Response**:
```json
{
  "message": "Successfully transferred asset shiva from sale to malli"
}
```

---

### 6. Read Asset

**Endpoint**: `GET /getasset/:id`

**Example**:
```bash
GET http://localhost:3000/adluru/shiva
```

---

### 7. Get Asset Transaction History

**Endpoint**: `GET /get/:id/history`

**Example**:
```bash
GET http://localhost:3000/adluru/shiva/history
```

---

This structured guide provides steps to set up the Hyperledger Fabric network, deploy the asset transfer chaincode, and interact with the blockchain using RESTful APIs through the Node.js application.
