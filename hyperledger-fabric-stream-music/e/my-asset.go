/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// MyAsset stores a value
type MyAsset struct {
	Key   string `json:"key"`
	Value *Dong
}
type Dong struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Token string `json:"token"`
}
