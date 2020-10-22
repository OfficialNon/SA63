// Code generated by entc, DO NOT EDIT.

package illness

const (
	// Label holds the string label denoting the illness type in the database.
	Label = "illness"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIllnessName holds the string denoting the illness_name field in the database.
	FieldIllnessName = "illness_name"

	// EdgeIllnessDiagnose holds the string denoting the illness_diagnose edge name in mutations.
	EdgeIllnessDiagnose = "illness_diagnose"

	// Table holds the table name of the illness in the database.
	Table = "illnesses"
	// IllnessDiagnoseTable is the table the holds the illness_diagnose relation/edge.
	IllnessDiagnoseTable = "diagnoses"
	// IllnessDiagnoseInverseTable is the table name for the Diagnose entity.
	// It exists in this package in order to avoid circular dependency with the "diagnose" package.
	IllnessDiagnoseInverseTable = "diagnoses"
	// IllnessDiagnoseColumn is the table column denoting the illness_diagnose relation/edge.
	IllnessDiagnoseColumn = "illness_illness_diagnose"
)

// Columns holds all SQL columns for illness fields.
var Columns = []string{
	FieldID,
	FieldIllnessName,
}

var (
	// IllnessNameValidator is a validator for the "Illness_Name" field. It is called by the builders before save.
	IllnessNameValidator func(string) error
)
