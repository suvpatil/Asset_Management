/*
Copyright IBM Corp. 2016 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
		 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	//"bytes"
	"strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"	
)

//var myLogger = logging.MustGetLogger("asset_mgm")

// AssetManagementChaincode is simple chaincode implementing a basic Asset Management system
// with access control enforcement at chaincode level.
// Look here for more information on how to implement access control at chaincode level:
// https://github.com/hyperledger/fabric/blob/master/docs/tech/application-ACL.md
// An asset is simply represented by a string.
type AssetManagementChaincode struct {
}

// Init method will be called during deployment.
// The deploy transaction metadata is supposed to contain the administrator cert
func (t *AssetManagementChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	//myLogger.Debug("Init Chaincode...")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	// Create ownership table
	err := stub.CreateTable("AssetsOwnership", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "traderLoginUserName", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "isBuyer", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "isSeller", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "selectedBuyerName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "purchaseOrder", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "totalPrice", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "currency", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "deliveryDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "incoterm", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "paymentConditions", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "articleId1", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "articleDesc1", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "articleQuantity1", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "articleId2", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "articleDesc2", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "articleQuantity2", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "buyerPaymentConfrimation", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "sellerInfoCounterParty", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "buyerBankCommitment", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "sellerForfaitInvoice", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "invoiceStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "paymentStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "contractStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "deliveryStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "isOrderConfirmed", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "deliveryTrackingId", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating AssetsOnwership table.")
	}

	
	//myLogger.Debug("Init Chaincode...done")

	return nil, nil
}

func (t *AssetManagementChaincode) assign(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var contractSlice []string
	var valSplit []string
	var results []string
	//myLogger.Debug("Assign...")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}
	

	contarctString := args[0]
	contractSlice = strings.Split(contarctString, ",")
	for i := range contractSlice{
		valSplit = strings.Split(contractSlice[i], ":")
		results = append(results, valSplit[1])
	}

	//asset := results[0]
	traderLoginUserName := results[0]
	isBuyer := results[1]
	selectedBuyerName := results[2]
	//owner := args[1]
	
	// Register assignment
	//myLogger.Debugf("New owner of [%s] is [%s]", asset, owner)

	ok, err := stub.InsertRow("AssetsOwnership", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: traderLoginUserName}},
			&shim.Column{Value: &shim.Column_String_{String_: isBuyer}},
			&shim.Column{Value: &shim.Column_String_{String_: isSeller}},
			&shim.Column{Value: &shim.Column_String_{String_: selectedBuyerName}},
			&shim.Column{Value: &shim.Column_String_{String_: purchaseOrder}},
			&shim.Column{Value: &shim.Column_String_{String_: totalPrice}},
			&shim.Column{Value: &shim.Column_String_{String_: currency}},
			&shim.Column{Value: &shim.Column_String_{String_: deliveryDate}},
			&shim.Column{Value: &shim.Column_String_{String_: incoterm}},
			&shim.Column{Value: &shim.Column_String_{String_: paymentConditions}},
			&shim.Column{Value: &shim.Column_String_{String_: articleId1}},
			&shim.Column{Value: &shim.Column_String_{String_: articleDesc1}},
			&shim.Column{Value: &shim.Column_String_{String_: articleQuantity1}},
			&shim.Column{Value: &shim.Column_String_{String_: articleId2}},
			&shim.Column{Value: &shim.Column_String_{String_: articleDesc2}},
			&shim.Column{Value: &shim.Column_String_{String_: articleQuantity2}},
			&shim.Column{Value: &shim.Column_String_{String_: buyerPaymentConfrimation}},
			&shim.Column{Value: &shim.Column_String_{String_: sellerInfoCounterParty}},
			&shim.Column{Value: &shim.Column_String_{String_: buyerBankCommitment}},
			&shim.Column{Value: &shim.Column_String_{String_: sellerForfaitInvoice}},
			&shim.Column{Value: &shim.Column_String_{String_: invoiceStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: paymentStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: contractStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: deliveryStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: isOrderConfirmed}},
			&shim.Column{Value: &shim.Column_String_{String_: deliveryTrackingId}},			
		},
	})

	if !ok && err == nil {
		return nil, errors.New("Asset was already assigned.")
	}

	//myLogger.Debug("Assign...done!")

	return nil, err
}


// Invoke will be called for every transaction.
// Supported functions are the following:
// "assign(asset, owner)": to assign ownership of assets. An asset can be owned by a single entity.
// Only an administrator can call this function.
// "transfer(asset, newOwner)": to transfer the ownership of an asset. Only the owner of the specific
// asset can call this function.
// An asset is any string to identify it. An owner is representated by one of his ECert/TCert.
func (t *AssetManagementChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Handle different functions
	if function == "assign" {
		// Assign ownership
		return t.assign(stub, args)
	
	}

	return nil, errors.New("Received unknown function invocation")
}

// Query callback representing the query of a chaincode
// Supported functions are the following:
// "query(asset)": returns the owner of the asset.
// Anyone can invoke this function.
func (t *AssetManagementChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	//myLogger.Debugf("Query [%s]", function)
	var jsonAsBytes []byte
	//var buffer bytes.Buffer	
	
	if function != "query" {
		return nil, errors.New("Invalid query function name. Expecting 'query' but found '" + function + "'")
	}

	var err error

	if len(args) != 1 {
		//myLogger.Debug("Incorrect number of arguments. Expecting name of an asset to query")
		return nil, errors.New("Incorrect number of arguments. Expecting name of an asset to query")
	}

	// Who is the owner of the asset?
	traderLoginUserName := args[0]

	///myLogger.Debugf("Arg [%s]", string(asset))

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: traderLoginUserName}}
	columns = append(columns, col1)

	row, err := stub.GetRow("AssetsOwnership", columns)
	if err != nil {
		//myLogger.Debugf("Failed retriving asset [%s]: [%s]", string(asset), err)
		return nil, fmt.Errorf("Failed retriving asset [%s]: [%s]", string(traderLoginUserName), err)
	}

	//myLogger.Debugf("Query done [% x]", row.Columns[1].GetBytes())
	//buffer.WriteString(row.Columns[0].GetString_())
	//buffer.WriteString(row.Columns[1].GetString_())
	//buffer.WriteString(row.Columns[2].GetString_())
	
	//row.Columns[0]
	//jsonAsBytes, _ = json.Marshal(buffer.String())
	jsonAsBytes, _ = json.Marshal(row)
	return jsonAsBytes, nil
}

func main() {
	
	err := shim.Start(new(AssetManagementChaincode))
	if err != nil {
		fmt.Printf("Error starting AssetManagementChaincode: %s", err)
	}
}	
