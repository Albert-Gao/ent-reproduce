// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"enttry/entgen/profile"
	"enttry/entgen/tenant"
	"enttry/entgen/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Profile is the model entity for the Profile schema.
type Profile struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// NickName holds the value of the "nick_name" field.
	NickName string `json:"nick_name,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID string `json:"owner_id,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID string `json:"tenant_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProfileQuery when eager-loading is set.
	Edges ProfileEdges `json:"edges"`
}

// ProfileEdges holds the relations/edges for other nodes in the graph.
type ProfileEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Tenant holds the value of the tenant edge.
	Tenant *Tenant `json:"tenant,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProfileEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// TenantOrErr returns the Tenant value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProfileEdges) TenantOrErr() (*Tenant, error) {
	if e.loadedTypes[1] {
		if e.Tenant == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tenant.Label}
		}
		return e.Tenant, nil
	}
	return nil, &NotLoadedError{edge: "tenant"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Profile) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case profile.FieldID, profile.FieldNickName, profile.FieldOwnerID, profile.FieldTenantID:
			values[i] = new(sql.NullString)
		case profile.FieldCreateTime, profile.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Profile", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Profile fields.
func (pr *Profile) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case profile.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pr.ID = value.String
			}
		case profile.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				pr.CreateTime = value.Time
			}
		case profile.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				pr.UpdateTime = value.Time
			}
		case profile.FieldNickName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nick_name", values[i])
			} else if value.Valid {
				pr.NickName = value.String
			}
		case profile.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				pr.OwnerID = value.String
			}
		case profile.FieldTenantID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				pr.TenantID = value.String
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Profile entity.
func (pr *Profile) QueryOwner() *UserQuery {
	return (&ProfileClient{config: pr.config}).QueryOwner(pr)
}

// QueryTenant queries the "tenant" edge of the Profile entity.
func (pr *Profile) QueryTenant() *TenantQuery {
	return (&ProfileClient{config: pr.config}).QueryTenant(pr)
}

// Update returns a builder for updating this Profile.
// Note that you need to call Profile.Unwrap() before calling this method if this Profile
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Profile) Update() *ProfileUpdateOne {
	return (&ProfileClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Profile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Profile) Unwrap() *Profile {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("entgen: Profile is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Profile) String() string {
	var builder strings.Builder
	builder.WriteString("Profile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("create_time=")
	builder.WriteString(pr.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(pr.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("nick_name=")
	builder.WriteString(pr.NickName)
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(pr.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("tenant_id=")
	builder.WriteString(pr.TenantID)
	builder.WriteByte(')')
	return builder.String()
}

// Profiles is a parsable slice of Profile.
type Profiles []*Profile

func (pr Profiles) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
