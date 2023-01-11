package ip

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
)

func TestGetInternalIP(t *testing.T) {
	ip, err := GetInternalIP()
	fmt.Println("internal ip: ", ip, err)
}
