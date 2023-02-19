package snowflake

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		dataCenterID int8
		machineId    int8
	}
	tests := []struct {
		name string
		args args
		want Snowflake
	}{
		{
			name: "Test 1",
			args: args{
				dataCenterID: 13,
				machineId:    18,
			},
		}, {
			name: "Test 2",
			args: args{
				dataCenterID: 12,
				machineId:    21,
			},
		}, {
			name: "Test 3",
			args: args{
				dataCenterID: 12,
				machineId:    18,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.dataCenterID, tt.args.machineId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}
