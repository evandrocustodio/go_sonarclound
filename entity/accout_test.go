package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "J@j")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func Test_AddAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "J@j")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(account.Client.Accounts))
}
