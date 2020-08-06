package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "SuperHero API")
	})
	router.HandleFunc("/all", all)

	router.HandleFunc("/heroes", heroes)

	router.HandleFunc("/villains", villains)

	router.HandleFunc("/search/{name}", searchName)

	router.HandleFunc("/id/{id}", byID).Methods("GET")

	router.HandleFunc("/id/{id}", removeByID).Methods("DELETE")

	router.HandleFunc("/new", newSuper).Methods("POST")

	fmt.Println("Servidor ok!")
	http.ListenAndServe(":8000", router)
}

type Super struct {
	Uuid         int    `json:"uuid"`
	Name         string `json:"name"`
	FullName     string `json:"fullName"`
	Intelligence string `json:"intelligence"`
	Power        string `json:"power"`
	Occupation   string `json:"occupation"`
	Image        string `json:"image"`
	Group        string `json:"group"`
	Relatives    string `json:"relatives"`
}

func all(w http.ResponseWriter, r *http.Request) {

	rows, err := ConnectDB().Query("SELECT * from super")

	if err != nil {
		fmt.Print(err)
	}

	searchSuper(rows, err, w)

}

func heroes(w http.ResponseWriter, r *http.Request) {

	rows, err := ConnectDB().Query("SELECT * from super where alignment = 'good'")

	if err != nil {
		fmt.Print(err)
	}

	searchSuper(rows, err, w)

}

func villains(w http.ResponseWriter, r *http.Request) {

	rows, err := ConnectDB().Query("SELECT * from super where alignment = 'bad'")

	if err != nil {
		fmt.Print(err)
	}

	searchSuper(rows, err, w)
}

func searchName(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	rows, err := ConnectDB().Query("SELECT * from super where UPPER(name) = UPPER('" + params["name"] + "')")

	if err != nil {
		fmt.Print(err)
	}

	searchSuper(rows, err, w)
}

func byID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	rows, err := ConnectDB().Query("SELECT * from super where uuid = '" + params["id"] + "'")

	if err != nil {
		fmt.Print(err)
	}

	searchSuper(rows, err, w)
}

func removeByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	fmt.Print(params["id"])

	_, err := ConnectDB().Query("DELETE FROM super WHERE uuid = '" + params["id"] + "'")

	if err != nil {
		fmt.Print(err)
	}

	fmt.Fprintln(w, "Registro excluido com sucesso!")
}

func newSuper(w http.ResponseWriter, r *http.Request) {

	var super Super
	_ = json.NewDecoder(r.Body).Decode(&super)

	if super.Name == "" {
		fmt.Fprintln(w, "Nome do super não pode ser nulo!")
	} else {
		rows, err := ConnectDB().Query("SELECT max(uuid) + 1 uuid FROM super")

		if err != nil {
			fmt.Print(err)
		}

		for rows.Next() {
			var uuid int
			err = rows.Scan(&uuid)
			if err != nil {
				fmt.Print(err)
			}
			super.Uuid = uuid
		}

		statement := "INSERT INTO super(uuid, name, full_name) VALUES($1, $2, $3)"

		stmt, err := ConnectDB().Exec(statement, super.Uuid, super.Name, super.FullName)

		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(stmt)
		fmt.Fprintln(w, "Super cadastrado com sucesso!")

	}

}

func searchSuper(rows *sql.Rows, err error, w http.ResponseWriter) {
	for rows.Next() {
		var uuid int
		var name string
		var fullName string
		var connections string
		var power string
		var occupation string
		var alignment string
		var image string

		err = rows.Scan(&uuid, &name, &fullName, &connections, &power, &occupation, &image, &alignment)
		if err != nil {
			fmt.Print(err)
		}

		var connectionsMap map[string]interface{}
		if err = json.Unmarshal([]byte(connections), &connectionsMap); err != nil {
			fmt.Println(err)
		}

		var powerMap map[string]interface{}
		if err = json.Unmarshal([]byte(power), &powerMap); err != nil {
			fmt.Println(err)
		}

		super := Super{
			Uuid:         uuid,
			Name:         name,
			FullName:     fullName,
			Intelligence: powerMap["intelligence"].(string),
			Power:        powerMap["power"].(string),
			Occupation:   occupation,
			Image:        image,
			Group:        connectionsMap["group-affiliation"].(string),
			Relatives:    connectionsMap["relatives"].(string),
		}

		json.NewEncoder(w).Encode(super)

	}
}

func ConnectDB() *sql.DB {

	connStr := "user=postgres dbname=projeto-go password=root host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print(err)
	}

	return db
}
