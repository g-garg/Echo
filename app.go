package main

import (
	"encoding/json"
	"log"
	"net/http"
    	"io/ioutil"
	"gopkg.in/mgo.v2/bson"
	. "myapp/config"
	. "myapp/dao"
	. "myapp/models"
	"github.com/labstack/echo/v4"
)

var config = Config{}
var dao = MoviesDAO{}

// // GET list of movies
// func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
// 	movies, err := dao.FindAll()
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJson(w, http.StatusOK, movies)
// }
//
// // GET a movie by its ID
// func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	movie, err := dao.FindById(params["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
// 		return
// 	}
// 	respondWithJson(w, http.StatusOK, movie)
// }
//
// // POST a new movie
// func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	var movie Movie
// 	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	movie.ID = bson.NewObjectId()
// 	if err := dao.Insert(movie); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJson(w, http.StatusCreated, movie)
// }
//
// // PUT update an existing movie
// func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	var movie Movie
// 	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	if err := dao.Update(movie); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
// }
//
// // DELETE an existing movie
// func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	var movie Movie
// 	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	if err := dao.Delete(movie); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
// }
//
// func respondWithError(w http.ResponseWriter, code int, msg string) {
// 	respondWithJson(w, code, map[string]string{"error": msg})
// }
//
// func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
// 	response, _ := json.Marshal(payload)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
// 	r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
// 	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
// 	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
// 	r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")
// 	if err := http.ListenAndServe(":3000", r); err != nil {
// 		log.Fatal(err)
// 	}
// }

func print(c echo.Context) error {
    resp, err := http.Get("https://data.covid19india.org/v4/min/data.min.json")
    if err != nil {
      log.Fatalln(err)
    }
     body, err := ioutil.ReadAll(resp.Body)
       if err != nil {
          log.Fatalln(err)
       }
     var data map[string]interface{}
     err1 := json.Unmarshal([]byte(string(body)), &data)
     if err1 != nil {
         log.Fatalln(err1)
     }
     for key, val := range data {
         state:= val.(map[string]interface{})
         result := state["total"].(map[string]interface{})
        confirmed := result["confirmed"].(float64)
        deceased := result["deceased"].(float64)
        recovered := result["recovered"].(float64)
//          log.Print(result["confirmed"])
//          log.Print(result["deceased"])
//          log.Print(result["recovered"])
         log.Print(key)
         log.Print(confirmed - deceased - recovered)

//      log.Print(data["AP"])
    //Convert the body to type string
//        sb := string(body)
// //        log.Printf(sb)
// 	   return c.String(http.StatusOK, sb)
        covid1 := Covid{ID: bson.NewObjectId(), State: key, PatientCount: confirmed - deceased - recovered}
//         covid1.ID := bson.NewObjectId()
// 	    covid1.state := key
// 	     covid1.PatientCount := confirmed - deceased - recovered

             // To get model's collection, just call to mgm.Coll() method.
//      err2 := mgm.Coll("covid").Insert(&covid)
  log.Print(covid1)
    err2 := dao.Insert(covid1)
     if err2 != nil {
             log.Fatalln(err2)
         }
    }

     return c.JSON(http.StatusOK, "done")
	}

func main() {
	e := echo.New()
    e.GET("/", print)
	e.Logger.Fatal(e.Start(":8185"))
}