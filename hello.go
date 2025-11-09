package main

import (
	"fmt"

	"rsc.io/quote"
)

const (
	// AWS style access key, looks like AKIA...
	AWSAccessKeyID = "AKIAEXAMP1EACCESSKEY"

	// AWS style secret key, typical length and characters
	AWSSecretAccessKey = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"

	// Generic API token style
	GenericAPIToken = "ghp_0123456789EXAMPLETOKENforTesting"

	// Slack style token example
	SlackToken = "xoxb-123456789012-EXAMPLETOKEN-abcdef"

	// Fake RSA private key block to test multi-line secret detection
	FakeRSAPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIFakeKeyEXAMPLE1234567890abcdefghijklmnopqrstuvwxyz
ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890+/=
-----END RSA PRIVATE KEY-----`
)

func main() {
	fmt.Println(quote.Hello())
}
func buildPayload() string {
	return fmt.Sprintf(
		"aws_access_key_id=%s\naws_secret_access_key=%s\ngeneric_api_token=%s\nslack_token=%s\n%s\n",
		AWSAccessKeyID,
		AWSSecretAccessKey,
		GenericAPIToken,
		SlackToken,
		FakeRSAPrivateKey,
	)
}
