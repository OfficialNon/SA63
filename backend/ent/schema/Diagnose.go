package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

type Diagnose struct {
	ent.Schema
}

func (Diagnose) Fields() []ent.Field {

	return []ent.Field{}
}
func (Diagnose) Edges() []ent.Edge {

	return []ent.Edge{
		edge.From("symptoms", Symptoms.Type).Ref("symptoms_diagnose").Unique(),
		edge.From("illness", Illness.Type).Ref("illness_diagnose").Unique(),
		edge.From("patient", Patient.Type).Ref("patient_diagnose").Unique(),
		edge.From("doctor", Doctor.Type).Ref("doctor_diagnose").Unique(),
	}

}
