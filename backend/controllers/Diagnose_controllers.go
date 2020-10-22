package controllers
 
import (
	"context"
	"strconv"
	"fmt"
	//"time"

	"github.com/gin-gonic/gin"
	"github.com/ADMIN/app/ent"
	"github.com/ADMIN/app/ent/doctor"
	"github.com/ADMIN/app/ent/patient"
	"github.com/ADMIN/app/ent/illness"
	"github.com/ADMIN/app/ent/symptoms"
	"github.com/ADMIN/app/ent/diagnose"
)
 
type DiagnoseController struct {
	client *ent.Client
	router gin.IRouter
}

type Diagnose struct {
	Doctor   	int
	Patient     int
	Illness 	int
	Symptoms    int
}
// CreateDiagnose handles POST requests for adding diagnose entities
// @Summary Create diagnose
// @Description Create diagnose
// @ID create-diagnose
// @Accept   json
// @Produce  json
// @Param diagnose body Diagnose true "Diagnose entity"
// @Success 200 {object} ent.Diagnose
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diagnoses [post]
func (ctl *DiagnoseController) CreateDiagnose(c *gin.Context) {
	obj := Diagnose{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Diagnose binding failed",
		})
		return
	}
	d, err := ctl.client.Doctor.
		Query().
		Where(doctor.IDEQ(int(obj.Doctor))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Doctor not found",
		})
		return
	}
	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patient))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Patient not found",
		})
		return
	}
	me, err := ctl.client.Illness.
		Query().
		Where(illness.IDEQ(int(obj.Illness))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Illness not found",
		})
		return
	}
	m, err := ctl.client.Symptoms.
		Query().
		Where(symptoms.IDEQ(int(obj.Symptoms))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Symptoms not found",
		})
		return
	}
	da, err := ctl.client.Diagnose.
		Create().
		SetDoctor(d).
		SetPatient(p).
		SetIllness(me).
		SetSymptoms(m).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": true,
		"data":   da,
	})
}
 // GetDiagnose handles GET requests to retrieve a diagnose entity
// @Summary Get a diagnose entity by ID
// @Description get diagnose by ID
// @ID get-drugAllergy
// @Produce  json
// @Param id path int true "Diagnose ID"
// @Success 200 {object} ent.Diagnose
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diagnoses/{id} [get]
func (ctl *DiagnoseController) GetDiagnose(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	b, err := ctl.client.Diagnose.
		Query().
		Where(diagnose.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, b)
}
 // ListDiagnose handles request to get a list of diagnose entities
// @Summary List Diagnose entities
// @Description list Diagnose entities
// @ID list-diagnose
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Diagnose
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diagnoses [get]
func (ctl *DiagnoseController) ListDiagnose(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {limit = int(limit64)}
	}
  
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {offset = int(offset64)}
	}
  
	diagnoses, err := ctl.client.Diagnose.
		Query().
		WithDoctor().
		WithPatient().
		WithIllness().
		WithSymptoms().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, diagnoses)
 }

 // DeleteDiagnose handles DELETE requests to delete a diagnose entity
// @Summary Delete a diagnose entity by ID
// @Description get diagnose by ID
// @ID delete-diagnose
// @Produce  json
// @Param id path int true "Diagnose ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diagnoses/{id} [delete]
func (ctl *DiagnoseController) DeleteDiagnose(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.Diagnose.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateDiagnose handles PUT requests to update a Diagnose entity
// @Summary Update a diagnose entity by ID
// @Description update diagnose by ID
// @ID update-diagnose
// @Accept   json
// @Produce  json
// @Param id path int true "Diagnose ID"
// @Param drugAllergy body ent.Diagnose true "Diagnose entity"
// @Success 200 {object} ent.Diagnose
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diagnoses/{id} [put]
func (ctl *DiagnoseController) UpdateDiagnose(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := ent.Diagnose{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "booking binding failed",
		})
		return
	}
	obj.ID = int(id)
	b, err := ctl.client.Diagnose.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}
	c.JSON(200, b)
} 
 
 // NewDiagnoseController creates and registers handles for the diagnose controller
func NewDiagnoseController(router gin.IRouter, client *ent.Client) *DiagnoseController {
	da := &DiagnoseController{
		client: client,
		router: router,
	}

	da.register()

	return da

}

func (ctl *DiagnoseController) register() {
	diagnoses := ctl.router.Group("/diagnoses")
	diagnoses.GET("", ctl.ListDiagnose)
	// CRUD
	diagnoses.POST("", ctl.CreateDiagnose)
	diagnoses.GET(":id", ctl.GetDiagnose)
	diagnoses.PUT(":id", ctl.UpdateDiagnose)
	diagnoses.DELETE(":id", ctl.DeleteDiagnose)
}
 
 