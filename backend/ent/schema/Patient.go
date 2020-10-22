package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Patient struct {
	ent.Schema
}

func (Patient) Fields() []ent.Field {

	return []ent.Field{

		field.String("Patient_Name").NotEmpty(),
	}
}
func (Patient) Edges() []ent.Edge {

	return []ent.Edge{
		edge.To("patient_diagnose", Diagnose.Type),
	}
}
