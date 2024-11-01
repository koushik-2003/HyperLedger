/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

 package chaincode

 import (
	 "encoding/json"
	 "fmt"
 
	 "github.com/hyperledger/fabric-contract-api-go/contractapi"
 )
 
 // SmartContract provides functions for managing an Asset
 type SmartContract struct {
	 contractapi.Contract
 }
 
 // Asset describes the basic details of what makes up an asset
 type Asset struct {
	 ID          string json:"ID"
	 DealerID    string json:"DEALERID"
	 MSISDN      string json:"MSISDN"
	 MPIN        string json:"MPIN"
	 Balance     int    json:"BALANCE"
	 Status      string json:"STATUS"
	 TransAmount int    json:"TRANSAMOUNT"
	 TransType   string json:"TRANSTYPE"
	 Remarks     string json:"REMARKS"
 }
 
 // InitLedger adds a base set of assets to the ledger
 func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	 assets := []Asset{
		 {ID: "asset1", DealerID: "dealer1", MSISDN: "1234567890", MPIN: "1234", Balance: 300, Status: "active", TransAmount: 0, TransType: "", Remarks: ""},
		 {ID: "asset2", DealerID: "dealer2", MSISDN: "0987654321", MPIN: "5678", Balance: 400, Status: "active", TransAmount: 0, TransType: "", Remarks: ""},
	 }
 
	 for _, asset := range assets {
		 assetJSON, err := json.Marshal(asset)
		 if err != nil {
			 return err
		 }
		 err = ctx.GetStub().PutState(asset.ID, assetJSON)
		 if err != nil {
			 return fmt.Errorf("failed to put to world state. %v", err)
		 }
	 }
	 return nil
 }
 
 // CreateAsset creates a new asset in the world state
 func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, id, dealerID, msisdn, mpin string, balance int, status string, transAmount int, transType, remarks string) error {
	 exists, err := s.AssetExists(ctx, id)
	 if err != nil {
		 return err
	 }
	 if exists {
		 return fmt.Errorf("the asset %s already exists", id)
	 }
 
	 asset := Asset{
		 ID:          id,
		 DealerID:    dealerID,
		 MSISDN:      msisdn,
		 MPIN:        mpin,
		 Balance:     balance,
		 Status:      status,
		 TransAmount: transAmount,
		 TransType:   transType,
		 Remarks:     remarks,
	 }
	 assetJSON, err := json.Marshal(asset)
	 if err != nil {
		 return err
	 }
	 return ctx.GetStub().PutState(id, assetJSON)
 }
 
 // ReadAsset retrieves an asset from the world state by ID
 func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	 assetJSON, err := ctx.GetStub().GetState(id)
	 if err != nil {
		 return nil, fmt.Errorf("failed to read from world state: %v", err)
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
 
 // UpdateAsset updates an existing asset in the world state
 func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id, dealerID, msisdn, mpin string, balance int, status string, transAmount int, transType, remarks string) error {
	 exists, err := s.AssetExists(ctx, id)
	 if err != nil {
		 return err
	 }
	 if !exists {
		 return fmt.Errorf("the asset %s does not exist", id)
	 }
 
	 asset := Asset{
		 ID:          id,
		 DealerID:    dealerID,
		 MSISDN:      msisdn,
		 MPIN:        mpin,
		 Balance:     balance,
		 Status:      status,
		 TransAmount: transAmount,
		 TransType:   transType,
		 Remarks:     remarks,
	 }
	 assetJSON, err := json.Marshal(asset)
	 if err != nil {
		 return err
	 }
 
	 return ctx.GetStub().PutState(id, assetJSON)
 }
 
 // DeleteAsset removes an asset from the world state
 func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	 exists, err := s.AssetExists(ctx, id)
	 if err != nil {
		 return err
	 }
	 if !exists {
		 return fmt.Errorf("the asset %s does not exist", id)
	 }
 
	 return ctx.GetStub().DelState(id)
 }
 
 // AssetExists checks if an asset with the given ID exists in the world state
 func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	 assetJSON, err := ctx.GetStub().GetState(id)
	 if err != nil {
		 return false, fmt.Errorf("failed to read from world state: %v", err)
	 }
	 return assetJSON != nil, nil
 }
 
 // TransferAsset updates the dealer ID of an asset and returns the previous dealer ID
 func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, id, newDealerID string) (string, error) {
	 asset, err := s.ReadAsset(ctx, id)
	 if err != nil {
		 return "", err
	 }
 
	 oldDealerID := asset.DealerID
	 asset.DealerID = newDealerID
 
	 assetJSON, err := json.Marshal(asset)
	 if err != nil {
		 return "", err
	 }
 
	 err = ctx.GetStub().PutState(id, assetJSON)
	 if err != nil {
		 return "", err
	 }
 
	 return oldDealerID, nil
 }
 
 // GetAllAssets retrieves all assets from the world state
 func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
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
 
 // GetAssetTransactionHistory retrieves the transaction history for an asset by ID
 func (s *SmartContract) GetAssetTransactionHistory(ctx contractapi.TransactionContextInterface, id string) ([]map[string]interface{}, error) {
	 resultsIterator, err := ctx.GetStub().GetHistoryForKey(id)
	 if err != nil {
		 return nil, err
	 }
	 defer resultsIterator.Close()
 
	 var records []map[string]interface{}
	 for resultsIterator.HasNext() {
		 queryResponse, err := resultsIterator.Next()
		 if err != nil {
			 return nil, err
		 }
 
		 record := map[string]interface{}{
			 "TxId":      queryResponse.TxId,
			 "Value":     string(queryResponse.Value),
			 "Timestamp": queryResponse.Timestamp,
			 "IsDelete":  queryResponse.IsDelete,
		 }
		 records = append(records, record)
	 }
	 return records, nil
 }