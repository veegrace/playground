package prototest

import (
	"fmt"

	"veegrace/playground/golang/prototest/pb"
)

// Important: this has to be in the same package as the generated protobuf code
func (e *pb.MyError) Error() string {
	return fmt.Sprintf("Error: %d: %s", e.Status, e.Reason)
}
