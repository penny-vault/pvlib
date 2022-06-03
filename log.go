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
	"encoding/hex"

	"github.com/rs/zerolog"
)

func (t *Transaction) MarshalZerologObject(e *zerolog.Event) {
	e.Str("ID", hex.EncodeToString(t.ID)).
		Str("CompositeFigi", t.CompositeFIGI).
		Str("Ticker", t.Ticker).
		Str("Kind", t.Kind).
		Time("Date", t.Date).
		Float64("TotalValue", t.TotalValue).
		Float64("PricePerShare", t.PricePerShare).
		Float64("Shares", t.Shares).
		Float64("Commission", t.Commission).
		Str("Source", t.Source).
		Str("SourceID", t.SourceID).
		Str("Memo", t.Memo).
		Strs("Tags", t.Tags).
		Str("TaxDisposition", t.TaxDisposition).
		Bool("Cleared", t.Cleared)
}
