/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	_ "bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// MyAssetContract contract for managing CRUD for MyAsset
type MyAssetContract struct {
	contractapi.Contract
}

// MyAssetExists returns true when asset with given ID exists in world state
func (c *MyAssetContract) ddongExists(ctx contractapi.TransactionContextInterface, key string) (bool, error) {
	data, err := ctx.GetStub().GetState(key)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// init
func (c *MyAssetContract) InitDDong(ctx contractapi.TransactionContextInterface, key string, dongid string, name string, token string) error {
	exists, err := c.ddongExists(ctx, key)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if exists {
		return fmt.Errorf("The asset %s already exists", key)
	}

	ddong := Dong{
		ID:    dongid,
		Name:  name,
		Token: token,
	}
	dongbyte, _ := json.Marshal(ddong)

	return ctx.GetStub().PutState(key, dongbyte)
}

// func (c *MyAssetContract) InitDDong(ctx contractapi.TransactionContextInterface, arg []string) error {
// 	exists, err := c.ddongExists(ctx, arg[0])
// 	if err != nil {
// 		return fmt.Errorf("Could not read from world state. %s", err)
// 	} else if exists {
// 		return fmt.Errorf("The asset %s already exists", arg[0])
// 	}

// 	ddong := Dong{
// 		ID:    arg[1],
// 		Name:  arg[2],
// 		Token: arg[3],
// 	}
// 	dongbyte, _ := json.Marshal(ddong)

// 	return ctx.GetStub().PutState(arg[0], dongbyte)
// }

// ReadMyAsset retrieves an instance of MyAsset from the world state
func (c *MyAssetContract) ReadMyAsset(ctx contractapi.TransactionContextInterface, key string) (*Dong, error) {
	exists, err := c.ddongExists(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("The asset %s does not exist", key)
	}

	bytes, _ := ctx.GetStub().GetState(key)

	dong := new(Dong)

	err = json.Unmarshal(bytes, dong)

	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal world state data to type MyAsset")
	}

	return dong, nil
}

// UpdateMyAsset retrieves an instance of MyAsset from the world state and updates its value
func (c *MyAssetContract) UpdateMyAsset(ctx contractapi.TransactionContextInterface, key string, newValue string) error {
	exists, err := c.ddongExists(ctx, key)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", key)
	}
	myAsset, err := c.ReadMyAsset(ctx, key)

	if err != nil {
		return err
	}
	myAsset.Name = newValue

	bytes, _ := json.Marshal(myAsset)

	return ctx.GetStub().PutState(key, bytes)
}

// DeleteMyAsset deletes an instance of MyAsset from the world state
func (c *MyAssetContract) DeleteMyAsset(ctx contractapi.TransactionContextInterface, key string) error {
	exists, err := c.ddongExists(ctx, key)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", key)
	}

	return ctx.GetStub().DelState(key)
}
