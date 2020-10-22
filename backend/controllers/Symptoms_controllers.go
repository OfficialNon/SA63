package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/ADMIN/app/ent"
   "github.com/ADMIN/app/ent/symptoms"
   "github.com/gin-gonic/gin"
)
 
// SymptomsController defines the struct for the symptoms controller
type SymptomsController struct {
   client *ent.Client
   router gin.IRouter
}
// CreateSymptoms handles POST requests for adding symptoms entities
// @Summary Create symptoms
// @Description Create symptoms
// @ID create-symptoms
// @Accept   json
// @Produce  json
// @Param symptoms body ent.Symptoms true "Symptoms entity"
// @Success 200 {object} ent.Symptoms
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptomss [post]
func (ctl *SymptomsController) CreateSymptoms(c *gin.Context) {
	obj := ent.Symptoms{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Symptoms binding failed",
		})
		return
	}
  
	me, err := ctl.client.Symptoms.
		Create().
		SetManner(obj.Manner).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, gin.H{
		"status" : true,
		"data" : me,
		"id" : me.ID, 
	})
 }
 // GetSymptoms handles GET requests to retrieve a symptoms entity
// @Summary Get a symptoms entity by ID
// @Description get symptoms by ID
// @ID get-symptoms
// @Produce  json
// @Param id path int true "Symptoms ID"
// @Success 200 {object} ent.Symptoms
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptomss/{id} [get]
func (ctl *SymptomsController) GetSymptoms(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	me, err := ctl.client.Symptoms.
		Query().
		Where(symptoms.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, me)
 }
 // ListSymptoms handles request to get a list of symptoms entities
// @Summary List symptoms entities
// @Description list symptoms entities
// @ID list-symptoms
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Symptoms
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptomss [get]
func (ctl *SymptomsController) ListSymptoms(c *gin.Context) {
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
  
	symptomss, err := ctl.client.Symptoms.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, symptomss)
 }
 // DeleteSymptoms handles DELETE requests to delete a symptoms entity
// @Summary Delete a symptoms entity by ID
// @Description get symptoms by ID
// @ID delete-symptoms
// @Produce  json
// @Param id path int true "Symptoms ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptomss/{id} [delete]
func (ctl *SymptomsController) DeleteSymptoms(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Symptoms.
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
 // UpdateSymptoms handles PUT requests to update a symptoms entity
// @Summary Update a symptoms entity by ID
// @Description update symptoms by ID
// @ID update-symptoms
// @Accept   json
// @Produce  json
// @Param id path int true "Symptoms ID"
// @Param symptoms body ent.Symptoms true "Symptoms entity"
// @Success 200 {object} ent.Symptoms
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptomss/{id} [put]
func (ctl *SymptomsController) UpdateSymptoms(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Symptoms{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Symptoms binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	p, err := ctl.client.Symptoms.
		UpdateOne(&obj).
		SetManner(obj.Manner).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, p)
 }
 // NewSymptomsController creates and registers handles for the symptoms controller
func NewSymptomsController(router gin.IRouter, client *ent.Client) *SymptomsController {
	uc := &SymptomsController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitSymptomsController registers routes to the main engine
 func (ctl *SymptomsController) register() {
	symptomss := ctl.router.Group("/symptomss")
  
	symptomss.GET("", ctl.ListSymptoms)
  
	// CRUD
	symptomss.POST("", ctl.CreateSymptoms)
	symptomss.GET(":id", ctl.GetSymptoms)
	symptomss.PUT(":id", ctl.UpdateSymptoms)
	symptomss.DELETE(":id", ctl.DeleteSymptoms)
 }
 
 