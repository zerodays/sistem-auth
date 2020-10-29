package permission

type Type struct {
	Code                string
	ImplicitPermissions []Type
}

var (
	UserRead = Type{
		Code:                "user:read",
		ImplicitPermissions: nil,
	}
	UserWrite = Type{
		Code:                "user:write",
		ImplicitPermissions: []Type{UserRead},
	}

	InventoryRead = Type{
		Code:                "inventory:read",
		ImplicitPermissions: nil,
	}
	InventoryWrite = Type{
		Code:                "inventory:write",
		ImplicitPermissions: []Type{InventoryRead},
	}
)

var All = []Type{
	UserRead,
	UserWrite,
	InventoryRead,
	InventoryWrite,
}
