package tests

import (
	"encoding/json"
	"github.com/zerodays/sistem-auth/permission"
	"testing"
)

func TestPermissionMarshalJSON(t *testing.T) {
	permissions := []permission.Type{
		permission.UserRead,
		permission.UserWrite,
		permission.InventoryRead,
		permission.InventoryWrite,
	}

	expected := []string{
		"\"user:read\"",
		"\"user:write\"",
		"\"inventory:read\"",
		"\"inventory:write\"",
	}

	for i, perm := range permissions {
		res, err := json.Marshal(perm)
		if err != nil {
			t.Error(err)
		}

		if string(res) != expected[i] {
			t.Errorf("Result for %v is incorrect. Expected \"%s\", got \"%s\"\n", perm, expected[i], string(res))
		}
	}
}

func equal(t1, t2 permission.Type) bool {
	if t1.Code != t2.Code {
		return false
	}

	if len(t1.ImplicitPermissions) != len(t2.ImplicitPermissions) {
		return false
	}

	for i := 0; i < len(t1.ImplicitPermissions); i++ {
		if !equal(t1.ImplicitPermissions[i], t2.ImplicitPermissions[i]) {
			return false
		}
	}

	return true
}

func TestPermissionUmnarshaslJSON(t *testing.T) {
	permissions := []string{
		"\"user:read\"",
		"\"user:write\"",
		"\"inventory:read\"",
		"\"inventory:write\"",
	}

	expected := []permission.Type{
		permission.UserRead,
		permission.UserWrite,
		permission.InventoryRead,
		permission.InventoryWrite,
	}

	for i, perm := range permissions {
		var res permission.Type
		err := json.Unmarshal([]byte(perm), &res)
		if err != nil {
			t.Error(err)
		}

		if !equal(expected[i], res) {
			t.Errorf("Result for \"%s\" is incorrect. Expected %v, got %v\n", perm, expected[i], res)
		}
	}
}
