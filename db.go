package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func DBInit() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}
	connStr := "user=postgres dbname=projeto-go password=root host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print(err)
	}
	//fmt.Print(db)

	var checkDatabase string
	db.QueryRow("SELECT to_regclass('public.super')").Scan(&checkDatabase)

	if checkDatabase == "" {
		fmt.Println("Database Created")
		createSQL := "CREATE TABLE public.super (uuid integer NOT NULL, name character varying(255) NOT NULL,full_name character varying(255)," +
			"connections character varying(1000), power character varying(1000), occupation character varying(1000)," +
			"image character varying(1000), alignment character varying(255),CONSTRAINT pk_id_super PRIMARY KEY (uuid));"
		db.Query(createSQL)
	}

	for i := 1; i < 730; i++ {
		resposta, err := cliente.Get("https://superheroapi.com/api/3194744643949892/" + strconv.Itoa(i))
		if err != nil {
			fmt.Println("[main] Erro ao abrir a pagina do Google Brasil. Erro: ", err.Error())
			return
		}
		defer resposta.Body.Close()

		if resposta.StatusCode == 200 {
			corpo, err := ioutil.ReadAll(resposta.Body)
			if err != nil {
				fmt.Println("[main] Erro ao ler o conteudo da pagina do Google Brasil. Erro: ", err.Error())
				return
			}

			super := make(map[string]interface{})
			err = json.Unmarshal(corpo, &super)
			if err != nil {
				fmt.Println(err)
				return
			}

			jsonConnections, _ := json.Marshal(super["connections"])
			jsonPowerstats, _ := json.Marshal(super["powerstats"])

			statement := "INSERT INTO super(uuid, name, full_name, connections, power, occupation, alignment, image) VALUES($1, $2, $3, $4, $5, $6, $7, $8)"

			stmt, err := db.Exec(statement, super["id"], super["name"], super["biography"].(map[string]interface{})["full-name"], string(jsonConnections), string(jsonPowerstats), super["work"].(map[string]interface{})["occupation"], super["biography"].(map[string]interface{})["alignment"], super["image"].(map[string]interface{})["url"])
			if err != nil {
				fmt.Print(err)
			}
			fmt.Println(stmt)

		}
	}

}
