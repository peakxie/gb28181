// -*-go-*-
// Copyright 2020 Justine Alexandra Roberts Tunney
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

package base 

import (
	"errors"
	"fmt"
)

%% machine www_authenticate;
%% include base "base.rl";
%% write data;

func ParseWwwAuth(data []byte) (params *Params, err error) {
	if data == nil {
		return nil, nil
	}
	params = new(Params)
	cs := 0
	p := 0
	pe := len(data)
	eof := len(data)
	buf := make([]byte, len(data))
	amt := 0
	mark := 0
	var name string

	%% main := WWWAuthenticateMessage;
	%% write init;
	%% write exec;

	if cs < www_authenticate_first_final {
		if p == pe {
			return nil, errors.New(fmt.Sprintf("Incomplete data: %s", data))
		} else {
			return nil, errors.New(fmt.Sprintf("Error in data at pos %d: %s", p, data))
		}
	}

	return params, nil
}


