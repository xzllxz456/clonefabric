/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const getStateError = "world state get error"

type MockStub struct {
	shim.ChaincodeStubInterface
	mock.Mock
}

func (ms *MockStub) GetState(key string) ([]byte, error) {
	args := ms.Called(key)

	return args.Get(0).([]byte), args.Error(1)
}

func (ms *MockStub) PutState(key string, value []byte) error {
	args := ms.Called(key, value)

	return args.Error(0)
}

func (ms *MockStub) DelState(key string) error {
	args := ms.Called(key)

	return args.Error(0)
}

type MockContext struct {
	contractapi.TransactionContextInterface
	mock.Mock
}

func (mc *MockContext) GetStub() shim.ChaincodeStubInterface {
	args := mc.Called()

	return args.Get(0).(*MockStub)
}

func configureStub() (*MockContext, *MockStub) {
	var nilBytes []byte

	testRr := new(Rr)
	testRr.Value = "set value"
	rrBytes, _ := json.Marshal(testRr)

	ms := new(MockStub)
	ms.On("GetState", "statebad").Return(nilBytes, errors.New(getStateError))
	ms.On("GetState", "missingkey").Return(nilBytes, nil)
	ms.On("GetState", "existingkey").Return([]byte("some value"), nil)
	ms.On("GetState", "rrkey").Return(rrBytes, nil)
	ms.On("PutState", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).Return(nil)
	ms.On("DelState", mock.AnythingOfType("string")).Return(nil)

	mc := new(MockContext)
	mc.On("GetStub").Return(ms)

	return mc, ms
}

func TestRrExists(t *testing.T) {
	var exists bool
	var err error

	ctx, _ := configureStub()
	c := new(RrContract)

	exists, err = c.RrExists(ctx, "statebad")
	assert.EqualError(t, err, getStateError)
	assert.False(t, exists, "should return false on error")

	exists, err = c.RrExists(ctx, "missingkey")
	assert.Nil(t, err, "should not return error when can read from world state but no value for key")
	assert.False(t, exists, "should return false when no value for key in world state")

	exists, err = c.RrExists(ctx, "existingkey")
	assert.Nil(t, err, "should not return error when can read from world state and value exists for key")
	assert.True(t, exists, "should return true when value for key in world state")
}

func TestCreateRr(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(RrContract)

	err = c.CreateRr(ctx, "statebad", "some value")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

	err = c.CreateRr(ctx, "existingkey", "some value")
	assert.EqualError(t, err, "The asset existingkey already exists", "should error when exists returns true")

	err = c.CreateRr(ctx, "missingkey", "some value")
	stub.AssertCalled(t, "PutState", "missingkey", []byte("{\"value\":\"some value\"}"))
}

func TestReadRr(t *testing.T) {
	var rr *Rr
	var err error

	ctx, _ := configureStub()
	c := new(RrContract)

	rr, err = c.ReadRr(ctx, "statebad")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when reading")
	assert.Nil(t, rr, "should not return Rr when exists errors when reading").
		rr, err = c.ReadRr(ctx, "missingkey")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when reading")
	assert.Nil(t, rr, "should not return Rr when key does not exist in world state when reading")

	rr, err = c.ReadRr(ctx, "existingkey")
	assert.EqualError(t, err, "Could not unmarshal world state data to type Rr", "should error when data in key is not Rr")
	assert.Nil(t, rr, "should not return Rr when data in key is not of type Rr")

	rr, err = c.ReadRr(ctx, "rrkey")
	expectedRr := new(Rr)
	expectedRr.Value = "set value"
	assert.Nil(t, err, "should not return error when Rr exists in world state when reading")
	assert.Equal(t, expectedRr, rr, "should return deserialized Rr from world state")
}

func TestUpdateRr(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(RrContract)

	err = c.UpdateRr(ctx, "statebad", "new value")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when updating")

	err = c.UpdateRr(ctx, "missingkey", "new value")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when updating")

	err = c.UpdateRr(ctx, "rrkey", "new value")
	expectedRr := new(Rr)
	expectedRr.Value = "new value"
	expectedRrBytes, _ := json.Marshal(expectedRr)
	assert.Nil(t, err, "should not return error when Rr exists in world state when updating")
	stub.AssertCalled(t, "PutState", "rrkey", expectedRrBytes)
}

func TestDeleteRr(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(RrContract)

	err = c.DeleteRr(ctx, "statebad")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

	err = c.DeleteRr(ctx, "missingkey")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when deleting")

	err = c.DeleteRr(ctx, "rrkey")
	assert.Nil(t, err, "should not return error when Rr exists in world state when deleting")
	stub.AssertCalled(t, "DelState", "rrkey")
}
