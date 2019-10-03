package handler

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"net/http"

	_ "github.com/amila-ku/botbackend/api/docs"
//	"github.com/amila-ku/botbackend/pkg/entity"
//	store "github.com/amila-ku/botbackend/pkg/store"
	"github.com/gorilla/mux"
	swagger "github.com/swaggo/http-swagger"
	"github.com/sirupsen/logrus"
)


// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ChatBot Server !")
	log.Info("Home page rendered")
}

// HandleRequests contains handler mappings
func HandleRequests() {

	// creates a new mux
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.PathPrefix("/swagger/").Handler(swagger.Handler(swagger.URL("http://localhost:10000/swagger/doc.json")))

	// chatbot needs webhook endpoint with GET and POST. GET endpoint is used
	router.HandleFunc("/webhook", verification).Methods("GET")
	router.HandleFunc("/webhook", handleMassages).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", router))
}

// Handle messages godoc
// @Summary Handle incoming messages
// @Description Handle incoming messages
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} entity.APIError "We need ID!!"
// @Failure 404 {object} entity.APIError "Can not find ID"
// @Failure 500 {object} entity.APIError "We had a problem"
// @Router /webhook [post]
func handleMassages(w http.ResponseWriter, r *http.Request) {
	
	log.Info("Start Handling Messages")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

	

	//prettyJSON(w, itm)

}

// Handle facebook challange  godoc
// @Summary Handle facebook challage verification
// @Description Handle facebook  challange verification
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} entity.APIError "We need ID!!"
// @Failure 404 {object} entity.APIError "Can not find ID"
// @Router /webhook [get]
func verification(w http.ResponseWriter, r *http.Request) {
	verifyToken := os.Getenv("VERIFY_TOKEN")

	challenge := r.URL.Query().Get("hub.challenge")
	mode := r.URL.Query().Get("hub.mode")
	token := r.URL.Query().Get("hub.verify_token")

	if mode != "" && token == verifyToken  {
		w.WriteHeader(200)
		w.Write([]byte(challenge))
		log.Info("Webhook verified")
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Error, wrong validation token"))
		log.Error("Webhook verification failed")
	}
}


func prettyJSON(w http.ResponseWriter, list interface{}) {
	pretty, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(pretty)
}
