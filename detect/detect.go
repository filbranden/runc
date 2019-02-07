package main

import (
	"fmt"

	"github.com/opencontainers/runc/libcontainer/cgroups/systemd"
)

func main() {
	cgroupSetup, err := systemd.DetectCgroupSetup()
	if err != nil {
		panic(fmt.Sprintf("%q", err))
	}
	fmt.Printf("%d\n", cgroupSetup)
}
