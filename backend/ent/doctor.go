// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/ADMIN/app/ent/doctor"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Doctor is the model entity for the Doctor schema.
type Doctor struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DoctorNAME holds the value of the "Doctor_NAME" field.
	DoctorNAME string `json:"Doctor_NAME,omitempty"`
	// DoctorEmail holds the value of the "Doctor_Email" field.
	DoctorEmail string `json:"Doctor_Email,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DoctorQuery when eager-loading is set.
	Edges DoctorEdges `json:"edges"`
}

// DoctorEdges holds the relations/edges for other nodes in the graph.
type DoctorEdges struct {
	// DoctorDiagnose holds the value of the doctor_diagnose edge.
	DoctorDiagnose []*Diagnose
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DoctorDiagnoseOrErr returns the DoctorDiagnose value or an error if the edge
// was not loaded in eager-loading.
func (e DoctorEdges) DoctorDiagnoseOrErr() ([]*Diagnose, error) {
	if e.loadedTypes[0] {
		return e.DoctorDiagnose, nil
	}
	return nil, &NotLoadedError{edge: "doctor_diagnose"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Doctor) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Doctor_NAME
		&sql.NullString{}, // Doctor_Email
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Doctor fields.
func (d *Doctor) assignValues(values ...interface{}) error {
	if m, n := len(values), len(doctor.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Doctor_NAME", values[0])
	} else if value.Valid {
		d.DoctorNAME = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Doctor_Email", values[1])
	} else if value.Valid {
		d.DoctorEmail = value.String
	}
	return nil
}

// QueryDoctorDiagnose queries the doctor_diagnose edge of the Doctor.
func (d *Doctor) QueryDoctorDiagnose() *DiagnoseQuery {
	return (&DoctorClient{config: d.config}).QueryDoctorDiagnose(d)
}

// Update returns a builder for updating this Doctor.
// Note that, you need to call Doctor.Unwrap() before calling this method, if this Doctor
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Doctor) Update() *DoctorUpdateOne {
	return (&DoctorClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Doctor) Unwrap() *Doctor {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Doctor is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Doctor) String() string {
	var builder strings.Builder
	builder.WriteString("Doctor(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", Doctor_NAME=")
	builder.WriteString(d.DoctorNAME)
	builder.WriteString(", Doctor_Email=")
	builder.WriteString(d.DoctorEmail)
	builder.WriteByte(')')
	return builder.String()
}

// Doctors is a parsable slice of Doctor.
type Doctors []*Doctor

func (d Doctors) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
