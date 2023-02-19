package snowflake

import (
	"fmt"
	"strings"
	"time"
)

type Snowflake struct {
	Value int64
}

func (s *Snowflake) String() string {
	var stringBuilder strings.Builder
	for i := uint(64); i >= 4; i -= 4 {
		if i == 48 || i == 32 {
			stringBuilder.WriteString("-")
		}
		stringBuilder.WriteString(fmt.Sprintf("%x", (s.Value>>i-4)&0xf))
	}
	stringBuilder.WriteString(fmt.Sprintf("%x", s.Value&0xf))
	return stringBuilder.String()
}

var lastSequence = 0

func New(dataCenterID, machineId int8) Snowflake {
	var snowflake Snowflake

	// Trim bits
	var timeStamp int64 = time.Now().UnixNano()
	timeStamp <<= 0x17 & 0x1ffffffffff

	// shift back
	timeStamp >>= 0x17 & 0x1ffffffffff

	// make sure datacenter ID 5 bits
	if dataCenterID > 0x1F {
		dataCenterID &= 0x1F
	}

	// make sure machineID is 5 bits
	if machineId > 0x1F {
		machineId &= 0x1F
	}

	lastSequence++

	// convert dataCenterID to 64 bit
	var datacenter64 int64 = int64(dataCenterID)
	var machineId64 int64 = int64(machineId)
	var lastSequence64 = int64(lastSequence)

	snowflake.Value = timeStamp << 22
	snowflake.Value |= datacenter64 << 17
	snowflake.Value |= machineId64 << 12
	snowflake.Value |= lastSequence64
	return snowflake
}
