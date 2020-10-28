package permission

type Permission string

const (
	UserRead       Permission = "user:read"
	UserWrite      Permission = "user:write"
	InventoryRead  Permission = "inventory:read"
	InventoryWrite Permission = "inventory:write"
)

var All = []Permission{
	UserRead,
	UserWrite,
	InventoryRead,
	InventoryWrite,
}
