// +build !pkcs11

package main

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/theupdateframework/notary"
	"github.com/theupdateframework/notary/passphrase"
)

func init() {
	NewNotaryCommand = func() *cobra.Command {
		commander := &notaryCommander{
			getRetriever: func() notary.PassRetriever { return passphrase.ConstantRetriever(testPassphrase) },
		}
		return commander.GetCommand()
	}
}

func rootOnHardware() bool {
	return false
}

// Per-test set up that is a no-op
func setUp(t *testing.T) {}

// no-op
func verifyRootKeyOnHardware(t *testing.T, rootKeyID string) {}
