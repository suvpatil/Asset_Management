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

type contract struct {	
	traderLoginUserName            string `json:"traderLoginUserName"`
	isBuyer                        string `json:"isBuyer"`     
	isSeller                       string `json:"isSeller"`
	selectedBuyerName              string `json:"selectedBuyerName"`
	purchaseOrder                  string `json:"purchaseOrder"`
	totalPrice                     string `json:"totalPrice"`
	currency                       string `json:"currency"`
	deliveryDate                   string `json:"deliveryDate"`
	incoterm                       string `json:"incoterm"`
	paymentConditions              string `json:"paymentConditions"`
	articleId1                     string `json:"articleId1"`
	articleDesc1                   string `json:"articleDesc1"`
	articleQuantity1               string `json:"articleQuantity1"`
	articleId2                     string `json:"articleId2"`
	articleDesc2                   string `json:"articleDesc2"`
	articleQuantity2               string `json:"articleQuantity2"`
	buyerPaymentConfrimation       string `json:"buyerPaymentConfrimation"`
	sellerInfoCounterParty         string `json:"sellerInfoCounterParty"`
	buyerBankCommitment            string `json:"buyerBankCommitment"`
	sellerForfaitInvoice           string `json:"sellerForfaitInvoice"`
	invoiceStatus                  string `json:"invoiceStatus"`
	paymentStatus                  string `json:"paymentStatus"`
	contractStatus                 string `json:"contractStatus"`
	deliveryStatus                 string `json:"deliveryStatus"`
	isOrderConfirmed               string `json:"isOrderConfirmed"`
	deliveryTrackingId             string `json:"deliveryTrackingId"`

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
	isSeller := results[2]
	selectedBuyerName := results[3]
	purchaseOrder := results[4]
	totalPrice := results[5]
	currency := results[6]
	deliveryDate := results[7]
	incoterm := results[8]
	paymentConditions := results[9]
	articleId1 := results[10]
	articleDesc1 := results[11]
	articleQuantity1 := results[12]
	articleId2 := results[13]
	articleDesc2 := results[14]
	articleQuantity2 := results[15]
	buyerPaymentConfrimation := results[16]
	sellerInfoCounterParty := results[17]
	buyerBankCommitment := results[18]
	sellerForfaitInvoice := results[19]
	invoiceStatus := results[20]
	paymentStatus := results[21]
	contractStatus := results[22]
	deliveryStatus := results[23]
	isOrderConfirmed := results[24]
	deliveryTrackingId := results[25]
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

func (t *AssetManagementChaincode) UpdateDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error

	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Need 6 arguments")
	}
	
	traderLoginUserName := args[0]
	purchaseOrder := args[1]
	//statusMessageToBeUodated: “paymentStatus”,
	//newStatusMessage: ””
	invoiceStatus := args[2]
	paymentStatus := args[3]
	contractStatus := args[4]
	deliveryStatus := args[5]

	ok, err := stub.ReplaceRow("AssetsOwnership", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: traderLoginUserName}},
			&shim.Column{Value: &shim.Column_String_{String_: purchaseOrder}},
			&shim.Column{Value: &shim.Column_String_{String_: invoiceStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: paymentStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: contractStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: deliveryStatus}},
		},
	})

	if !ok && err == nil {
		return nil, errors.New("Error in adding record.")
	}
	return nil, nil
}


func (t *AssetManagementChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Handle different functions
	if function == "assign" {		
		return t.assign(stub, args)	
	}else if function == "UpdateDetails" {		
		return t.UpdateDetails(stub, args)
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
	//resultArray := make(map[string]string)
	var contObj contract
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
	
	seller_name := row.Columns[0].GetString_()
	buyer_name := row.Columns[3].GetString_()
	if traderLoginUserName == seller_name || traderLoginUserName == buyer_name {
		contObj.traderLoginUserName = row.Columns[0].GetString_()
		contObj.isBuyer = row.Columns[1].GetString_()
		contObj.isSeller = row.Columns[2].GetString_()
		contObj.selectedBuyerName = row.Columns[3].GetString_()
		contObj.purchaseOrder = row.Columns[4].GetString_()
		contObj.totalPrice = row.Columns[5].GetString_()
		contObj.currency = row.Columns[6].GetString_()
		contObj.deliveryDate = row.Columns[7].GetString_()
		contObj.incoterm = row.Columns[8].GetString_()
		contObj.paymentConditions = row.Columns[9].GetString_()
		contObj.articleId1 = row.Columns[10].GetString_()
		contObj.articleDesc1 = row.Columns[11].GetString_()
		contObj.articleQuantity1 = row.Columns[12].GetString_()
		contObj.articleId2 = row.Columns[13].GetString_()
		contObj.articleDesc2 = row.Columns[14].GetString_()
		contObj.articleQuantity2 = row.Columns[15].GetString_()
		contObj.buyerPaymentConfrimation = row.Columns[16].GetString_()
		contObj.sellerInfoCounterParty = row.Columns[17].GetString_()
		contObj.buyerBankCommitment = row.Columns[18].GetString_()
		contObj.sellerForfaitInvoice = row.Columns[19].GetString_()
		contObj.invoiceStatus = row.Columns[20].GetString_()
		contObj.paymentStatus = row.Columns[21].GetString_()
		contObj.contractStatus = row.Columns[22].GetString_()
		contObj.deliveryStatus = row.Columns[23].GetString_()
		contObj.isOrderConfirmed = row.Columns[24].GetString_()
		contObj.deliveryTrackingId = row.Columns[25].GetString_()
		jsonAsBytes, _ = json.Marshal(contObj)
		return jsonAsBytes, nil
	}else {
		return nil
	}
	//myLogger.Debugf("Query done [% x]", row.Columns[1].GetBytes())
	//buffer.WriteString(row.Columns[0].GetString_())
	//buffer.WriteString(row.Columns[1].GetString_())
	//buffer.WriteString(row.Columns[2].GetString_())
	
	//row.Columns[0]
	//jsonAsBytes, _ = json.Marshal(buffer.String())
	//jsonAsBytes, _ = json.Marshal(row)
	
}

func main() {
	
	err := shim.Start(new(AssetManagementChaincode))
	if err != nil {
		fmt.Printf("Error starting AssetManagementChaincode: %s", err)
	}
}	
