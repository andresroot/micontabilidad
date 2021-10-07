package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var db *gorm.DB
var err error

type Tarea struct {
	Id          uint
	Titulo      string
	Descripcion string
	Estado      string
}
type Resultado struct {
	Resultado bool
}

func favicon(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s\n", r.RequestURI)
	w.Header().Set("Content-Type", "image/x-icon")
	w.Header().Set("Cache-Control", "public, max-age=7776000")
	fmt.Fprintln(w, "data:image/x-icon;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQEAYAAABPYyMiAAAABmJLR0T///////8JWPfcAAAACXBIWXMAAABIAAAASABGyWs+AAAAF0lEQVRIx2NgGAWjYBSMglEwCkbBSAcACBAAAeaR9cIAAAAASUVORK5CYII=\n")
}

var Raiz string

func main() {

	dsn := "host=192.168.1.7 user=postgres password=Murc4505 dbname=Base2020 port=5432 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}
	r := mux.NewRouter()

	Raiz, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	if Raiz == "C:\\micontabilidad\\bin" {
		Raiz = ""
	} else {
		Raiz = Raiz + "/"
	}

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir(Raiz+"static/"))))

	// default
	//r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	ts, err := template.ParseFiles(Raiz + "vista/login.html")
	//	if err != nil {
	//		log.Println(err.Error())
	//		http.Error(w, "Internal Server Error1111"+err.Error()+Raiz, 500)
	//		return
	//	}
	//	err = ts.Execute(w, nil)
	//	if err != nil {
	//		log.Println(err.Error())
	//		http.Error(w, "Internal Server Error222", 500)
	//	}
	//})

	// default
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var tareas []Tarea
		//db.Order("age desc, name").Find(&users)
		db.Order("Id").Find(&tareas)
		for _, miFila := range tareas {
			// bla bla
			log.Println(miFila.Titulo)
		}
		ts, err := template.ParseFiles(Raiz + "vista/menu.html")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error1111"+err.Error()+Raiz, 500)
			return
		}

		parametros := map[string]interface{}{
			"tareas": tareas,
		}

		err = ts.Execute(w, parametros)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error222", 500)
		}
	})

	r.Path("/Nuevatarea").HandlerFunc(Nuevatarea).Name("DocumentoNuevo")
	r.Path("/Vertarea/{codigo}").HandlerFunc(Vertarea).Name("DocumentoNuevo")
	r.Path("/Eliminartarea/{codigo}").HandlerFunc(Eliminartarea).Name("DocumentoNuevo")
	r.Path("/Editartarea").HandlerFunc(Editartarea).Name("DocumentoNuevo")

	http.ListenAndServe(":9990", r)
}

func Editartarea(w http.ResponseWriter, r *http.Request) {
	var tempVenta Tarea
	dsn := "host=192.168.1.7 user=postgres password=Murc4505 dbname=Base2020 port=5432 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la venta
	err = json.Unmarshal(b, &tempVenta)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	user := Tarea{0, tempVenta.Titulo, tempVenta.Descripcion, tempVenta.Estado}
	//var result bool
	db.Model(&user).Where("Id= ?", tempVenta.Id).Update("Titulo", tempVenta.Titulo).Update("Descripcion", tempVenta.Descripcion).Update("Estado", tempVenta.Estado)
	//db.Save(&user) // pass pointer of data to Create
	//db.Model(&product).Update("Price", 2000)
	var resultado bool
	resultado = true

	js, err := json.Marshal(Resultado{resultado})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func Nuevatarea(w http.ResponseWriter, r *http.Request) {
	var tempVenta Tarea
	dsn := "host=192.168.1.7 user=postgres password=Murc4505 dbname=Base2020 port=5432 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la venta
	err = json.Unmarshal(b, &tempVenta)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	user := Tarea{0, tempVenta.Titulo, tempVenta.Descripcion, tempVenta.Estado}
	//var result bool

	result := db.Create(&user) // pass pointer of data to Create
	if result.RowsAffected > 0 {
		var resultado bool
		resultado = true

		js, err := json.Marshal(Resultado{resultado})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	} else {
		var resultado bool
		resultado = false

		js, err := json.Marshal(Resultado{resultado})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

}

func Vertarea(w http.ResponseWriter, r *http.Request) {

	dsn := "host=192.168.1.7 user=postgres password=Murc4505 dbname=Base2020 port=5432 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	Codigo := mux.Vars(r)["codigo"]

	var miTarea Tarea
	db.Table("tareas").Where("Id = ?", Codigo).Scan(&miTarea)

	js, err := json.Marshal(miTarea)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	//result.RowsAffected
}

func Eliminartarea(w http.ResponseWriter, r *http.Request) {

	dsn := "host=192.168.1.7 user=postgres password=Murc4505 dbname=Base2020 port=5432 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	Codigo := mux.Vars(r)["codigo"]

	var miTarea Tarea
	db.Table("tareas").Where("Id = ?", Codigo).Scan(&miTarea)
	db.Delete(&miTarea, Codigo)

	var resultado bool
	resultado = true

	js, err := json.Marshal(Resultado{resultado})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	//result.RowsAffected
}
