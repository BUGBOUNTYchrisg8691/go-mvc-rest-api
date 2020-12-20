package bootstrap

import (
	"github.com/bugbountychris8691/go-mvc-rest-api/src/webapp/controllers"
	"net/http"
)

func BootApplication() {
	http.HandleFunc("/employees", controllers.GetEmployeeById)

	if err := http.ListenAndServe(":2019", nil); err != nil {
		panic(err)
	}
}