## Snowflake Algorithm

This is a Go package for generating Snowflake IDs. Snowflake is a distributed unique ID generator that is used in a variety of systems to generate unique IDs in a distributed environment. The Snowflake algorithm generates 64-bit IDs that are made up of a timestamp, a datacenter ID, a machine ID, and a sequence number.

The Snowflake ID consists of the following components:

| Timestamp | Datacenter ID | Machine  ID | Sequence Number |
|-----------|---------------|-------------|-----------------|
| 41 bits   | 5 bits        | 5 bits      | 12 bits         |

## Usage

The package provides a `Snowflake` struct that represents a generated Snowflake ID. The `New` function can be used to generate a new ID:

```go
func New(dataCenterID, machineID int8) Snowflake
```

Here is an example on how to use it 
```go
import (
    "github.com/oyamo/snowflake"
)
//..

s := snowflake.New(1, 1)
fmt.Println(s.String())

```

## LICENCE
MIT Licence
