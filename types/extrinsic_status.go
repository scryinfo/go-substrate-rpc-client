// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/scryinfo/go-substrate-rpc-client/scale"
)

// ExtrinsicStatus is an enum containing the result of an extrinsic submission
type ExtrinsicStatus struct {
	IsFuture    bool // 00:: Future
	IsReady     bool // 1:: Ready
	IsFinalized bool // 2:: Finalized(Hash)
	AsFinalized Hash
	IsUsurped   bool // 3:: Usurped(Hash)
	AsUsurped   Hash
	IsBroadcast bool // 4:: Broadcast(Vec<Text>)
	AsBroadcast []Text
	IsDropped   bool // 5:: Dropped
	IsInvalid   bool // 6:: Invalid
}

func (e *ExtrinsicStatus) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	switch b {
	case 0:
		e.IsFuture = true
	case 1:
		e.IsReady = true
	case 2:
		e.IsFinalized = true
		err = decoder.Decode(&e.AsFinalized)
	case 3:
		e.IsUsurped = true
		err = decoder.Decode(&e.AsUsurped)
	case 4:
		e.IsBroadcast = true
		err = decoder.Decode(&e.AsBroadcast)
	case 5:
		e.IsDropped = true
	case 6:
		e.IsInvalid = true
	}

	if err != nil {
		return err
	}

	return nil
}

func (e ExtrinsicStatus) Encode(encoder scale.Encoder) error {
	var err1, err2 error
	switch {
	case e.IsFuture:
		err1 = encoder.PushByte(0)
	case e.IsReady:
		err1 = encoder.PushByte(1)
	case e.IsFinalized:
		err1 = encoder.PushByte(2)
		err2 = encoder.Encode(e.AsFinalized)
	case e.IsUsurped:
		err1 = encoder.PushByte(3)
		err2 = encoder.Encode(e.AsUsurped)
	case e.IsBroadcast:
		err1 = encoder.PushByte(4)
		err2 = encoder.Encode(e.AsBroadcast)
	case e.IsDropped:
		err1 = encoder.PushByte(5)
	case e.IsInvalid:
		err1 = encoder.PushByte(6)
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}

func (e *ExtrinsicStatus) UnmarshalJSON(b []byte) error {
	input := strings.TrimSpace(string(b))
	if len(input) >= 2 && input[0] == '"' && input[len(input)-1] == '"' {
		input = input[1 : len(input)-1]
	}

	switch {
	case input == "future":
		e.IsFuture = true
		return nil
	case input == "ready":
		e.IsReady = true
		return nil
	case input == "dropped":
		e.IsDropped = true
		return nil
	case input == "invalid":
		e.IsInvalid = true
		return nil
	}

	// no simple case, decode into helper
	var tmp struct {
		AsFinalized Hash   `json:"finalized"`
		AsUsurped   Hash   `json:"usurped"`
		AsBroadcast []Text `json:"broadcast"`
	}
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	switch {
	case strings.HasPrefix(input, "{\"finalized\""):
		e.IsFinalized = true
		e.AsFinalized = tmp.AsFinalized
		return nil
	case strings.HasPrefix(input, "{\"usurped\""):
		e.IsUsurped = true
		e.AsUsurped = tmp.AsUsurped
		return nil
	case strings.HasPrefix(input, "{\"broadcast\""):
		e.IsBroadcast = true
		e.AsBroadcast = tmp.AsBroadcast
		return nil
	}

	return fmt.Errorf("unexpected JSON for ExtrinsicStatus, got %v", string(b))
}

func (e ExtrinsicStatus) MarshalJSON() ([]byte, error) {
	switch {
	case e.IsFuture:
		return []byte("\"future\""), nil
	case e.IsReady:
		return []byte("\"ready\""), nil
	case e.IsDropped:
		return []byte("\"dropped\""), nil
	case e.IsInvalid:
		return []byte("\"invalid\""), nil
	case e.IsFinalized:
		var tmp struct {
			AsFinalized Hash `json:"finalized"`
		}
		tmp.AsFinalized = e.AsFinalized
		return json.Marshal(tmp)
	case e.IsUsurped:
		var tmp struct {
			AsUsurped Hash `json:"usurped"`
		}
		tmp.AsUsurped = e.AsUsurped
		return json.Marshal(tmp)
	case e.IsBroadcast:
		var tmp struct {
			AsBroadcast []Text `json:"broadcast"`
		}
		tmp.AsBroadcast = e.AsBroadcast
		return json.Marshal(tmp)
	}
	return nil, fmt.Errorf("cannot marshal ExtrinsicStatus, got %#v", e)
}
