package db

import (
	"fmt"
	"testing"
)

func TestAbleToCreateAndDelete(t *testing.T) {
	sm, err := InitializeCassandra()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	secretId, err := sm.CreateSecret(0, "acoolsecret")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	secretContents, err := sm.GetSecret(secretId)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	if secretContents.EncryptedSecret != "acoolsecret" {
		t.FailNow()
	}
}
