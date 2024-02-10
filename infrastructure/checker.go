package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	brevo "github.com/getbrevo/brevo-go/lib"
)

const waitTimeSec = 10

func sendMail(instance string) {
	var ctx context.Context
	cfg := brevo.NewConfiguration()
	cfg.AddDefaultHeader("api-key", os.Getenv("BREVO_API_KEY"))

	br := brevo.NewAPIClient(cfg)
	_, _, err := br.AccountApi.GetAccount(ctx)
	if err != nil {
		fmt.Println("Error when calling AccountApi->get_account: ", err.Error())
		return
	}

	mail, _, err := br.TransactionalEmailsApi.SendTransacEmail(
		ctx,
		brevo.SendSmtpEmail{
			Sender: &brevo.SendSmtpEmailSender{Name: os.Getenv("BREVO_SENDER_NAME"), Email: os.Getenv("BREVO_SENDER_MAIL")},
			To: []brevo.SendSmtpEmailTo{
				{Email: os.Getenv("BREVO_SENDER_MAIL")},
			},
			Subject:     instance + " created!",
			TextContent: "Done.",
		},
	)

	if err != nil {
		fmt.Println("Error when trying to send an email: ", err.Error())
		return
	}
	fmt.Printf("Mail send successfully: %v", mail)
}

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
					sendMail(name)
					return
				}
				fmt.Printf("Waiting for %d secs before applying again.\n", waitTimeSec)
				time.Sleep(waitTimeSec * time.Second)
			}
		}
	}
}
