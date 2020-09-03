// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/crowdsecurity/crowdsec/cmd/api/ent/alert"
	"github.com/crowdsecurity/crowdsec/cmd/api/ent/blocker"
	"github.com/crowdsecurity/crowdsec/cmd/api/ent/decision"
	"github.com/crowdsecurity/crowdsec/cmd/api/ent/event"
	"github.com/crowdsecurity/crowdsec/cmd/api/ent/machine"
	"github.com/crowdsecurity/crowdsec/cmd/api/ent/meta"
	"github.com/crowdsecurity/crowdsec/cmd/api/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	alertFields := schema.Alert{}.Fields()
	_ = alertFields
	// alertDescCreatedAt is the schema descriptor for created_at field.
	alertDescCreatedAt := alertFields[0].Descriptor()
	// alert.DefaultCreatedAt holds the default value on creation for the created_at field.
	alert.DefaultCreatedAt = alertDescCreatedAt.Default.(func() time.Time)
	// alertDescUpdatedAt is the schema descriptor for updated_at field.
	alertDescUpdatedAt := alertFields[1].Descriptor()
	// alert.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	alert.DefaultUpdatedAt = alertDescUpdatedAt.Default.(func() time.Time)
	blockerFields := schema.Blocker{}.Fields()
	_ = blockerFields
	// blockerDescCreatedAt is the schema descriptor for created_at field.
	blockerDescCreatedAt := blockerFields[0].Descriptor()
	// blocker.DefaultCreatedAt holds the default value on creation for the created_at field.
	blocker.DefaultCreatedAt = blockerDescCreatedAt.Default.(func() time.Time)
	// blockerDescUpdatedAt is the schema descriptor for updated_at field.
	blockerDescUpdatedAt := blockerFields[1].Descriptor()
	// blocker.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	blocker.DefaultUpdatedAt = blockerDescUpdatedAt.Default.(func() time.Time)
	decisionFields := schema.Decision{}.Fields()
	_ = decisionFields
	// decisionDescCreatedAt is the schema descriptor for created_at field.
	decisionDescCreatedAt := decisionFields[0].Descriptor()
	// decision.DefaultCreatedAt holds the default value on creation for the created_at field.
	decision.DefaultCreatedAt = decisionDescCreatedAt.Default.(func() time.Time)
	// decisionDescUpdatedAt is the schema descriptor for updated_at field.
	decisionDescUpdatedAt := decisionFields[1].Descriptor()
	// decision.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	decision.DefaultUpdatedAt = decisionDescUpdatedAt.Default.(func() time.Time)
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescCreatedAt is the schema descriptor for created_at field.
	eventDescCreatedAt := eventFields[0].Descriptor()
	// event.DefaultCreatedAt holds the default value on creation for the created_at field.
	event.DefaultCreatedAt = eventDescCreatedAt.Default.(func() time.Time)
	// eventDescUpdatedAt is the schema descriptor for updated_at field.
	eventDescUpdatedAt := eventFields[1].Descriptor()
	// event.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	event.DefaultUpdatedAt = eventDescUpdatedAt.Default.(func() time.Time)
	machineFields := schema.Machine{}.Fields()
	_ = machineFields
	// machineDescCreatedAt is the schema descriptor for created_at field.
	machineDescCreatedAt := machineFields[0].Descriptor()
	// machine.DefaultCreatedAt holds the default value on creation for the created_at field.
	machine.DefaultCreatedAt = machineDescCreatedAt.Default.(func() time.Time)
	// machineDescUpdatedAt is the schema descriptor for updated_at field.
	machineDescUpdatedAt := machineFields[1].Descriptor()
	// machine.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	machine.DefaultUpdatedAt = machineDescUpdatedAt.Default.(func() time.Time)
	// machineDescIsValidated is the schema descriptor for isValidated field.
	machineDescIsValidated := machineFields[5].Descriptor()
	// machine.DefaultIsValidated holds the default value on creation for the isValidated field.
	machine.DefaultIsValidated = machineDescIsValidated.Default.(bool)
	metaFields := schema.Meta{}.Fields()
	_ = metaFields
	// metaDescCreatedAt is the schema descriptor for created_at field.
	metaDescCreatedAt := metaFields[0].Descriptor()
	// meta.DefaultCreatedAt holds the default value on creation for the created_at field.
	meta.DefaultCreatedAt = metaDescCreatedAt.Default.(func() time.Time)
	// metaDescUpdatedAt is the schema descriptor for updated_at field.
	metaDescUpdatedAt := metaFields[1].Descriptor()
	// meta.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	meta.DefaultUpdatedAt = metaDescUpdatedAt.Default.(func() time.Time)
}
