package consullook

import (
	"fmt"

	"code.byted.org/gopkg/consul"
)

func getServiceAddr(addr string) (string, error) {
	eds, err := consul.Lookup(addr)

	if err != nil {
		return "", err
	}

	if len(eds) == 0 {
		return "", fmt.Errorf("service discovery failed, cannot found job manager address")
	}

	return eds.GetOne().Addr, nil
}
