// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeProfiles holds the string denoting the profiles edge name in mutations.
	EdgeProfiles = "profiles"
	// EdgeTenants holds the string denoting the tenants edge name in mutations.
	EdgeTenants = "tenants"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ProfilesTable is the table that holds the profiles relation/edge.
	ProfilesTable = "profiles"
	// ProfilesInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ProfilesInverseTable = "profiles"
	// ProfilesColumn is the table column denoting the profiles relation/edge.
	ProfilesColumn = "owner_id"
	// TenantsTable is the table that holds the tenants relation/edge. The primary key declared below.
	TenantsTable = "user_tenants"
	// TenantsInverseTable is the table name for the Tenant entity.
	// It exists in this package in order to avoid circular dependency with the "tenant" package.
	TenantsInverseTable = "tenants"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
}

var (
	// TenantsPrimaryKey and TenantsColumn2 are the table columns denoting the
	// primary key for the tenants relation (M2M).
	TenantsPrimaryKey = []string{"user_id", "tenant_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)
