package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//_ "github.com/lib/pq"
)

/*const (
	//details about the db
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Aspergillus@123"
	dbname   = "first_db"
)*/

type registration struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var registrations = []registration{
	{ID: "1", Name: "Priyanshu", Email: "qwer@gmail.com"},
	{ID: "2", Name: "Aditya", Email: "zxcvb@gmail.com"},
	{ID: "3", Name: "Nipun", Email: "hello@gmail.com"},
}

func getRegistrations(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, registrations) //converting the above data structure into JSON

}

func addRegistrations(context *gin.Context) { //get the data from client in JSON format and it should be inside the request body
	var newRegistration registration
	if err := context.BindJSON(&newRegistration); err != nil {
		return
	}
	registrations = append(registrations, newRegistration)    //adding todos
	context.IndentedJSON(http.StatusCreated, newRegistration) //returning new Registrations
}

func main() {

	//connection string
	//psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//open database
	//db, err := sql.Open("postgres", psqlconn)
	//CheckError(err)

	//close database
	//defer db.Close()

	// Hardcoding the data
	//insertStmt := `insert into "Employee"("Name","EmpId") values('Priyanshu',21)`
	//_, e := db.Exec(insertStmt)
	//CheckError(e)

	router := gin.Default() //created a server
	router.GET("/registrations", getRegistrations)
	router.POST("/registrations", addRegistrations)

	router.Run("localhost:9090") //starting the server,specify the path inside it

}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
