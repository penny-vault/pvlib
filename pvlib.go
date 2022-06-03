/*
Copyright 2022

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

package pvlib

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type CompareOperator string

const (
	Equal       CompareOperator = "eq"
	GreaterThan CompareOperator = "gt"
	LessThan    CompareOperator = "lt"
	Like        CompareOperator = "like"
)

const (
	SellTransaction     = "SELL"
	BuyTransaction      = "BUY"
	DividendTransaction = "DIVIDEND"
	SplitTransaction    = "SPLIT"
	DepositTransaction  = "DEPOSIT"
	WithdrawTransaction = "WITHDRAW"
	InterestTransaction = "INTEREST"
)

type SearchCriteria struct {
	Name        *string
	StringValue *string
	IntValue    *int
	Comparator  CompareOperator
}

var apiKey string = "<not-set>"

func Authorize(myApiKey string) {
	apiKey = myApiKey
}

func Get(id []byte) (*Portfolio, error) {
	return nil, errors.New("not implemented")
}

// UploadTransactions uploads the provided transactions to the
// portfolio
func (p *Portfolio) UploadTransactions(transactions []*Transaction) error {
	log.Info().Msg("not implemented")
	return errors.New("not implemented")
}

// DownloadOrders retrieves the list of orders that should be executed on the account
func (p *Portfolio) DownloadOrders(transactions []*Transaction) error {
	log.Info().Msg("not implemented")
	return errors.New("not implemented")
}
