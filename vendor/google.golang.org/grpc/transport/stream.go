// +build !js !wasm

/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package transport

import (
	"io"
)

// Read reads all p bytes from the wire for this stream.
func (s *Stream) Read(p []byte) (n int, err error) {
	// Don't request a read if there was an error earlier
	if er := s.trReader.(*transportReader).er; er != nil {
		return 0, er
	}
	s.requestRead(len(p))
	return io.ReadFull(s.trReader, p)
}
