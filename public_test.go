package ip

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
)

func TestGetPublicIP(t *testing.T) {
	ip, err := GetPublicIP()
	fmt.Println("public ip: ", ip, err)
}
