package controller

import (
	"encoding/json"
	"net/http"

	"github.com/qlu1990/gos"
	"github.com/qlu1990/gos-example/model"
)

//AddPerson person
func AddPerson(c *gos.Context) {
	var person = model.Person{}
	err := c.PostBodyBind(&person)
	if err != nil {
		gos.Error(err)
		http.Error(c.ResponseWriter, "error json ", http.StatusNotFound)
		return
	}
	err = person.Add()
	if err != nil {
		gos.Error(err)
		http.Error(c.ResponseWriter, "fail add person ", http.StatusNotFound)
		return
	}
	gos.Response(c.ResponseWriter, "success", http.StatusOK)

}

// ListPersons list all person
func ListPersons(c *gos.Context) {
	persons, err := model.List()
	if err != nil {
		gos.Error(err)
		http.Error(c.ResponseWriter, "404 NotFound ", http.StatusNotFound)
		return
	}
	personsJSON, err := json.Marshal(persons)
	if err != nil {
		gos.Error(err)
		http.Error(c.ResponseWriter, "404 NotFound ", http.StatusNotFound)
		return
	}
	gos.Response(c.ResponseWriter, string(personsJSON), http.StatusOK)
}

// GetPerson get person by name
func GetPerson(c *gos.Context) {
	name := c.URIParam("name")
	persons, err := model.GetPersonByName(name)
	if err != nil {
		gos.Error(err)
		http.Error(c.ResponseWriter, "404 NotFound ", http.StatusNotFound)
		return
	}
	personsJSON, err := json.Marshal(persons)
	if err != nil {
		gos.Error(err)
		http.Error(c.ResponseWriter, "404 NotFound ", http.StatusNotFound)
		return
	}
	gos.Response(c.ResponseWriter, string(personsJSON), http.StatusOK)
}
