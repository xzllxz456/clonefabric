/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// RrContract contract for managing CRUD for Rr
type RrContract struct {
	contractapi.Contract
}

// RrExists returns true when asset with given ID exists in world state
func (c *RrContract) RrExists(ctx contractapi.TransactionContextInterface, rrID string) (bool, error) {
	data, err := ctx.GetStub().GetState(rrID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateRr creates a new instance of Rr
func (c *RrContract) CreateRr(ctx contractapi.TransactionContextInterface, rrID string, value string) error {
	exists, err := c.RrExists(ctx, rrID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if exists {
		return fmt.Errorf("The asset %s already exists", rrID)
	}

	rr := new(Rr)
	rr.Value = value

	bytes, _ := json.Marshal(rr)

	return ctx.GetStub().PutState(rrID, bytes)
}

// ReadRr retrieves an instance of Rr from the world state
func (c *RrContract) ReadRr(ctx contractapi.TransactionContextInterface, rrID string) (*Rr, error) {
	exists, err := c.RrExists(ctx, rrID)
	if err != nil {
		return nil, fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("The asset %s does not exist", rrID)
	}

	bytes, _ := ctx.GetStub().GetState(rrID)

	rr := new(Rr)

	err = json.Unmarshal(bytes, rr)

	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal world state data to type Rr")
	}

	return rr, nil
}

// UpdateRr retrieves an instance of Rr from the world state and updates its value
func (c *RrContract) UpdateRr(ctx contractapi.TransactionContextInterface, rrID string, newValue string) error {
	exists, err := c.RrExists(ctx, rrID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", rrID)
	}

	rr := new(Rr)
	rr.Value = newValue

	bytes, _ := json.Marshal(rr)

	return ctx.GetStub().PutState(rrID, bytes)
}

// DeleteRr deletes an instance of Rr from the world state
func (c *RrContract) DeleteRr(ctx contractapi.TransactionContextInterface, rrID string) error {
	exists, err := c.RrExists(ctx, rrID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", rrID)
	}

	return ctx.GetStub().DelState(rrID)
}
