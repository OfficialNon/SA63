package schema

import (
	"github.com/facebookincubator/ent"

	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Illness struct {
	ent.Schema
}

func (Illness) Fields() []ent.Field {

	return []ent.Field{

		field.String("Illness_Name").NotEmpty(),
	}

}
func (Illness) Edges() []ent.Edge {

	return []ent.Edge{
		edge.To("illness_diagnose", Diagnose.Type),
	}
}
