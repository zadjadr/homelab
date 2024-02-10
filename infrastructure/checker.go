package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const waitTimeSec = 10

func applyTerraform(az, name string) bool {
	fmt.Printf("Checking for %s in %s\n", name, az)
	cmd := exec.Command("terraform", "apply", "-auto-approve", "-var", "availability_domain="+az, "-var", "instance_name="+name)
	cmd.Stdout = nil
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err == nil
}

func main() {
	ociZones := os.Getenv("OCI_AZONES")
	if ociZones == "" {
		ociZones = "ldMg:EU-FRANKFURT-1-AD-1 ldMg:EU-FRANKFURT-1-AD-2 ldMg:EU-FRANKFURT-1-AD-3"
	}
	allAZ := strings.Split(ociZones, " ")

	ociInstanceNames := os.Getenv("OCI_INSTANCE_NAMES")
	if ociInstanceNames == "" {
		ociInstanceNames = "3.cloud 4.cloud"
	}
	allNames := strings.Split(ociInstanceNames, " ")

	for {
		for _, name := range allNames {
			for _, az := range allAZ {
				if applyTerraform(az, name) {
					fmt.Println("Succeeded!")
					return
				}
				fmt.Printf("Waiting for %d secs before applying again.\n", waitTimeSec)
				time.Sleep(waitTimeSec * time.Second)
			}
		}
	}
}
