package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/ADMIN/app/ent"
   "github.com/ADMIN/app/ent/illness"
   "github.com/gin-gonic/gin"
)
 
// IllnessController defines the struct for the illness controller
type IllnessController struct {
   client *ent.Client
   router gin.IRouter
}
// CreateIllness handles POST requests for adding illness entities
// @Summary Create illness
// @Description Create illness
// @ID create-illness
// @Accept   json
// @Produce  json
// @Param illness body ent.Illness true "Illness entity"
// @Success 200 {object} ent.Illness
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /illnesss [post]
func (ctl *IllnessController) CreateIllness(c *gin.Context) {
	obj := ent.Illness{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Illness binding failed",
		})
		return
	}
  
	me, err := ctl.client.Illness.
		Create().
		SetIllnessName(obj.IllnessName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, me)
 }
 // GetIllness handles GET requests to retrieve a illness entity
// @Summary Get a illness entity by ID
// @Description get illness by ID
// @ID get-illness
// @Produce  json
// @Param id path int true "Illness ID"
// @Success 200 {object} ent.Illness
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /illnesss/{id} [get]
func (ctl *IllnessController) GetIllness(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	me, err := ctl.client.Illness.
		Query().
		Where(illness.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, me)
 }
 // ListIllness handles request to get a list of illness entities
// @Summary List illness entities
// @Description list illness entities
// @ID list-illness
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Illness
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /illnesss [get]
func (ctl *IllnessController) ListIllness(c *gin.Context) {
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
  
	illnesss, err := ctl.client.Illness.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, illnesss)
 }
 // DeleteIllness handles DELETE requests to delete a illness entity
// @Summary Delete a illness entity by ID
// @Description get illness by ID
// @ID delete-illness
// @Produce  json
// @Param id path int true "Illness ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /illnesss/{id} [delete]
func (ctl *IllnessController) DeleteIllness(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Illness.
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
 // UpdateIllness handles PUT requests to update a illness entity
// @Summary Update a illness entity by ID
// @Description update illness by ID
// @ID update-illness
// @Accept   json
// @Produce  json
// @Param id path int true "Illness ID"
// @Param illness body ent.Illness true "Illness entity"
// @Success 200 {object} ent.Illness
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /illnesss/{id} [put]
func (ctl *IllnessController) UpdateIllness(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Illness{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Illness binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	p, err := ctl.client.Illness.
		UpdateOne(&obj).
		SetIllnessName(obj.IllnessName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, p)
 }
 // NewIllnessController creates and registers handles for the illness controller
func NewIllnessController(router gin.IRouter, client *ent.Client) *IllnessController {
	uc := &IllnessController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitIllnessController registers routes to the main engine
 func (ctl *IllnessController) register() {
	illnesss := ctl.router.Group("/illnesss")
  
	illnesss.GET("", ctl.ListIllness)
  
	// CRUD
	illnesss.POST("", ctl.CreateIllness)
	illnesss.GET(":id", ctl.GetIllness)
	illnesss.PUT(":id", ctl.UpdateIllness)
	illnesss.DELETE(":id", ctl.DeleteIllness)
 }
 
 