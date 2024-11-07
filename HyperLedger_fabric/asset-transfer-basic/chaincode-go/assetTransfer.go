/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

 package main

 import (
	 "encoding/json"
	 "fmt"
	 
	 "time"
 
	 "github.com/hyperledger/fabric-contract-api-go/contractapi"
 )
 

 type AssetTransfer struct {
	 contractapi.Contract
 }
 
 
 type Asset struct {
	 ID         string  `json:"ID"`
	 DealerID   string  `json:"DEALERID"`
	 MSISDN     string  `json:"MSISDN"`
	 MPIN       string  `json:"MPIN"`
	 Balance    float64 `json:"BALANCE"`
	 Status     string  `json:"STATUS"`
	 TransAmount float64 `json:"TRANSAMOUNT"`
	 TransType  string  `json:"TRANSTYPE"`
	 Remarks    string  `json:"REMARKS"`
 }
 
 
 func (s *AssetTransfer) InitLedger(ctx contractapi.TransactionContextInterface) error {
	 assets := []Asset{
		 {ID: "adluru", DealerID: "koushik", MSISDN: "741", MPIN: "2001", Balance: 300, Status: "active", TransAmount: 0, TransType: "", Remarks: ""},
		 {ID: "ram", DealerID: "surya", MSISDN: "963333", MPIN: "3002", Balance: 400, Status: "active", TransAmount: 0, TransType: "", Remarks: ""},
	 }
 
	 for _, asset := range assets {
		 assetJSON, err := json.Marshal(asset)
		 if err != nil {
			 return fmt.Errorf("failed to marshal asset: %v", err)
		 }
		 err = ctx.GetStub().PutState(asset.ID, assetJSON)
		 if err != nil {
			 return fmt.Errorf("failed to put asset to world state: %v", err)
		 }
	 }
 
	 return nil
 }
 
 
 func (s *AssetTransfer) CreateAsset(ctx contractapi.TransactionContextInterface, id, dealerId, msisdn, mpin string, balance float64, status string, transAmount float64, transType, remarks string) error {
	 exists, err := s.AssetExists(ctx, id)
	 if err != nil {
		 return err
	 }
	 if exists {
		 return fmt.Errorf("the asset %s already exists", id)
	 }
 
	 asset := Asset{
		 ID:         id,
		 DealerID:   dealerId,
		 MSISDN:     msisdn,
		 MPIN:       mpin,
		 Balance:    balance,
		 Status:     status,
		 TransAmount: transAmount,
		 TransType:  transType,
		 Remarks:    remarks,
	 }
 
	 assetJSON, err := json.Marshal(asset)
	 if err != nil {
		 return err
	 }
 
	 return ctx.GetStub().PutState(id, assetJSON)
 }
 
 
 func (s *AssetTransfer) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	 assetJSON, err := ctx.GetStub().GetState(id)
	 if err != nil {
		 return nil, fmt.Errorf("failed to read asset %s: %v", id, err)
	 }
	 if assetJSON == nil {
		 return nil, fmt.Errorf("the asset %s does not exist", id)
	 }
 
	 var asset Asset
	 err = json.Unmarshal(assetJSON, &asset)
	 if err != nil {
		 return nil, err
	 }
 
	 return &asset, nil
 }
 
 
 func (s *AssetTransfer) UpdateAsset(ctx contractapi.TransactionContextInterface, id, dealerId, msisdn, mpin string, balance float64, status string, transAmount float64, transType, remarks string) error {
	 exists, err := s.AssetExists(ctx, id)
	 if err != nil {
		 return err
	 }
	 if !exists {
		 return fmt.Errorf("the asset %s does not exist", id)
	 }
 
	 asset := Asset{
		 ID:         id,
		 DealerID:   dealerId,
		 MSISDN:     msisdn,
		 MPIN:       mpin,
		 Balance:    balance,
		 Status:     status,
		 TransAmount: transAmount,
		 TransType:  transType,
		 Remarks:    remarks,
	 }
 
	 assetJSON, err := json.Marshal(asset)
	 if err != nil {
		 return err
	 }
 
	 return ctx.GetStub().PutState(id, assetJSON)
 }
 
 
 func (s *AssetTransfer) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	 exists, err := s.AssetExists(ctx, id)
	 if err != nil {
		 return err
	 }
	 if !exists {
		 return fmt.Errorf("the asset %s does not exist", id)
	 }
 
	 return ctx.GetStub().DelState(id)
 }
 
 
 func (s *AssetTransfer) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	 assetJSON, err := ctx.GetStub().GetState(id)
	 if err != nil {
		 return false, err
	 }
 
	 return assetJSON != nil, nil
 }
 
 
 func (s *AssetTransfer) TransferAsset(ctx contractapi.TransactionContextInterface, id, newOwner string) (string, error) {
	 asset, err := s.ReadAsset(ctx, id)
	 if err != nil {
		 return "", err
	 }
 
	 oldOwner := asset.DealerID
	 asset.DealerID = newOwner
 
	 assetJSON, err := json.Marshal(asset)
	 if err != nil {
		 return "", err
	 }
 
	 err = ctx.GetStub().PutState(id, assetJSON)
	 if err != nil {
		 return "", err
	 }
 
	 return oldOwner, nil
 }
 

 func (s *AssetTransfer) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	 resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	 if err != nil {
		 return nil, err
	 }
	 defer resultsIterator.Close()
 
	 var assets []*Asset
	 for resultsIterator.HasNext() {
		 queryResponse, err := resultsIterator.Next()
		 if err != nil {
			 return nil, err
		 }
 
		 var asset Asset
		 err = json.Unmarshal(queryResponse.Value, &asset)
		 if err != nil {
			 return nil, err
		 }
		 assets = append(assets, &asset)
	 }
 
	 return assets, nil
 }
 

 func (s *AssetTransfer) GetAssetTransactionHistory(ctx contractapi.TransactionContextInterface, id string) ([]map[string]interface{}, error) {
	 resultsIterator, err := ctx.GetStub().GetHistoryForKey(id)
	 if err != nil {
		 return nil, err
	 }
	 defer resultsIterator.Close()
 
	 var records []map[string]interface{}
	 for resultsIterator.HasNext() {
		 response, err := resultsIterator.Next()
		 if err != nil {
			 return nil, err
		 }
 
		 record := map[string]interface{}{
			 "TxId":      response.TxId,
			 "Value":     string(response.Value),
			 "Timestamp": time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)),
			 "IsDelete":  response.IsDelete,
		 }
		 records = append(records, record)
	 }
 
	 return records, nil
 }
 
 func main() {
	 chaincode, err := contractapi.NewChaincode(&AssetTransfer{})
	 if err != nil {
		 fmt.Printf("Error creating asset-transfer chaincode: %v", err)
		 return
	 }
 
	 if err := chaincode.Start(); err != nil {
		 fmt.Printf("Error starting asset-transfer chaincode: %v", err)
	 }
 }
 
