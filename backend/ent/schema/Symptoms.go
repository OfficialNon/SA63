package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Symptoms struct {
	ent.Schema
}

func (Symptoms) Fields() []ent.Field {

	return []ent.Field{

		field.String("Manner").NotEmpty(),
		
	}
}
func (Symptoms) Edges() []ent.Edge {

	return []ent.Edge{
		edge.To("symptoms_diagnose", Diagnose.Type),
	}
}
