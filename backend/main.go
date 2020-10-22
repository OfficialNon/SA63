package main

import (
	"context"
	"log"

	"github.com/ADMIN/app/controllers"
	_ "github.com/ADMIN/app/docs"
	"github.com/ADMIN/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Doctors struct {
	Doctor []Doctor
}
type Doctor struct {
	Email string
	Name  string
}

type Patients struct {
	Patient []Patient
}
type Patient struct {
	Name string
}

type Illnesss struct {
	Illness []Illness
}
type Illness struct {
	Name string
}

type Symptomss struct {
	Symptoms []Symptoms
}
type Symptoms struct {
	Manner string
}

// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())
	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	v1 := router.Group("/api/v1")
	controllers.NewPatientController(v1, client)
	controllers.NewDoctorController(v1, client)
	controllers.NewIllnessController(v1, client)
	controllers.NewSymptomsController(v1, client)
	controllers.NewDiagnoseController(v1, client)
	// Set Doctor Data
	doctors := Doctors{
		Doctor: []Doctor{
			Doctor{"ADMIN@gmail.com", "ADMIN"},
			Doctor{"Ploymanee@gmail.com", "Ploymanee Chuyat"},
		},
	}
	for _, d := range doctors.Doctor {
		client.Doctor.
			Create().
			SetDoctorEmail(d.Email).
			SetDoctorNAME(d.Name).
			Save(context.Background())
	}
	// Set Illness Data
	illnesss := []string{"โรคเบาหวาน", "โรคอ้วน", "โรคกระเพาะอาหาร", "โรคหัวใจ"}
	for _, i := range illnesss {
		client.Illness.
			Create().
			SetIllnessName(i).
			Save(context.Background())
	}
	// Set Patient Data
	patients := []string{"สมหมาย", "สมชาย", "สมถุย", "สมศักดิ์"}
	for _, p := range patients {
		client.Patient.
			Create().
			SetPatientName(p).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
