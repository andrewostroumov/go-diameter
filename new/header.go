// Copyright 2013-2014 go-diameter authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package diam

import (
	"encoding/binary"
	"fmt"
)

// Diameter Header.
type Header struct {
	Version       uint8
	MessageLength uint32
	CommandFlags  uint8
	CommandCode   uint32
	ApplicationId uint32
	HopByHopId    uint32
	EndToEndId    uint32
}

const HeaderLength = 20 // Diameter header length.

func decodeHeader(data []byte) (*Header, error) {
	p := new(Header)
	if err := p.DecodeFromBytes(data); err != nil {
		return nil, err
	}
	return p, nil
}

// DecodeFromBytes decodes the bytes of a Diameter Header.
// It requires a dictionary Parser to be able to decode the Header
// by consulting the
func (h *Header) DecodeFromBytes(data []byte) error {
	if n := len(data); n < HeaderLength {
		return fmt.Errorf("Not enough data to decode Header: %d", n)
	}
	h.Version = data[0]
	h.MessageLength = uint24to32(data[1:4])
	h.CommandFlags = data[4]
	h.CommandCode = uint24to32(data[5:8])
	h.ApplicationId = binary.BigEndian.Uint32(data[8:12])
	h.HopByHopId = binary.BigEndian.Uint32(data[12:16])
	h.EndToEndId = binary.BigEndian.Uint32(data[16:20])
	return nil
}

func (h *Header) Serialize() ([]byte, error) {
	b := make([]byte, 20)
	b[0] = h.Version
	copy(b[1:], uint32to24(h.MessageLength))
	b[4] = h.CommandFlags
	copy(b[5:], uint32to24(h.CommandCode))
	binary.BigEndian.PutUint32(b[8:], h.ApplicationId)
	binary.BigEndian.PutUint32(b[12:], h.HopByHopId)
	binary.BigEndian.PutUint32(b[16:], h.EndToEndId)
	return b, nil
}

func (h *Header) String() string {
	return fmt.Sprintf("{Version:0x%x,Length:%d,Flags:0x%x,Code:%d,ApplicationId:%d,HopByHopId:0x%x,EndToEndId:0x%x}",
		h.Version,
		h.MessageLength,
		h.CommandFlags,
		h.CommandCode,
		h.ApplicationId,
		h.HopByHopId,
		h.EndToEndId,
	)
}