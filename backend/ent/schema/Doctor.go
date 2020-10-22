package schema

import (
	"github.com/facebookincubator/ent"

	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Doctor struct {
	ent.Schema
}

func (Doctor) Fields() []ent.Field {

	return []ent.Field{

		field.String("Doctor_NAME").NotEmpty(),

		field.String("Doctor_Email").NotEmpty(),
	}
}
func (Doctor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("doctor_diagnose", Diagnose.Type),
	}
}
