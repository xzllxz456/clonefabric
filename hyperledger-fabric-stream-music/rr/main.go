/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

func main() {
	rrContract := new(RrContract)
	rrContract.Info.Version = "0.0.1"
	rrContract.Info.Description = "My Smart Contract"
	rrContract.Info.License = new(metadata.LicenseMetadata)
	rrContract.Info.License.Name = "Apache-2.0"
	rrContract.Info.Contact = new(metadata.ContactMetadata)
	rrContract.Info.Contact.Name = "John Doe"

	chaincode, err := contractapi.NewChaincode(rrContract)
	chaincode.Info.Title = "rr chaincode"
	chaincode.Info.Version = "0.0.1"

	if err != nil {
		panic("Could not create chaincode from RrContract." + err.Error())
	}

	err = chaincode.Start()

	if err != nil {
		panic("Failed to start chaincode. " + err.Error())
	}
}
