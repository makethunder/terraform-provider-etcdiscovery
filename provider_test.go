package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

func TestProvider(test *testing.T) {
	provider := Provider().(*schema.Provider)
	if err := provider.InternalValidate(); err != nil {
		test.Fatalf("err: %s", err)
	}
}
