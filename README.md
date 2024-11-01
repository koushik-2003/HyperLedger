If you'd like to test all the API endpoints using Postman, here are examples for each API request you can use with the provided base URL. Follow the steps below to make the API calls.
Steps to Install and Setup the Network
Navigate to Your Project Directory:
cd HyperLedger_Fabric_Projet/HyperLedger_fabric Install Hyperledger Fabric:

curl -sSL https://bit.ly/2ysbOFE | bash -s -- 2.5.0

Go to the Fabric Test Network Directory:
cd fabric-samples/test-network

Stop Any Existing Network:
./network.sh down

Start the Network with CA Enabled and Create a Channel:
sudo ./network.sh up createChannel -ca -c mychannel

Deploy the Chaincode:
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-javascript -ccl javascript

Set Environment Variables for the Peer:
export PATH=${PWD}/../bin:$PATH export FABRIC_CFG_PATH=$PWD/../config/ export CORE_PEER_TLS_ENABLED=true export CORE_PEER_LOCALMSPID="Org1MSP" export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp export CORE_PEER_ADDRESS=localhost:7051

Invoke the Chaincode:
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}' Running the Node.js Application

Go to the Node.js Application Directory:
cd asset-transfer-basic/application-gateway-javascript

Install Required Dependencies:
npm i

Start the Node.js Server:
npm start

The server should start at http://localhost:3000.
Accessing and Testing the APIs Once the server is running locally, you can access the API endpoints using the following examples:

1. Initialize Ledger
POST /ledger/init

Example:
POST http://localhost:3000/ledger Response:

{ "message": "Ledger initialized successfully" }

2. Get All Assets
GET /getassets

Example:
GET http://localhost:3000/getassets

Response: A list of assets in the ledger.
3. Create Asset (Example for Dealer 3)
POST /createasset

Example:
{ ID: 'adluru', DEALERID: 'koushik', MSISDN: '741', MPIN: '2001', BALANCE: 300, STATUS: 'active', TRANSAMOUNT: 0, TRANSTYPE: '', REMARKS: '', } Response:

{ "message": "Asset adluru created successfully" }

4. Update Asset
PUT /updateasset

Example:
{ "id": "shiva", "dealerId": "sale", "msisdn": "258", "mpin": "123", "balance": 3625, "status": "active", "transAmount": 60, "transType": "debit", "remarks": "update balance" } Response:

{ "message": "Asset shiva updated successfully" }

5. Transfer Asset
POST /asset/transfer

Example:
{ "id": "shiva", "newOwner": "malli" } Response:

{ "message": "Successfully transferred asset shiva from sale to malli" }

6. Read Asset
GET /getasset/:id

Example:
GET http://localhost:3000/asset/asset1

7. Get Asset Transaction History
GET /get/:id/history

Example:
GET http://localhost:3000/adluru/shiva/history

This guide will help you set up and run your Hyperledger Fabric network, deploy the asset transfer chaincode, and use a Node.js application to interact with the blockchain via RESTful APIs.
