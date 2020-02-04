package core

import (
	// "encoding/json"
	// "reflect"	
	"fmt"
	"os"
	"flag"
	"log"
	"strconv"
	"net/http"
	"html/template"		

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)


type WebServerConfig struct{
	Port int
	Address string
	Config string
}

type Config struct {

}

func loadDefaultWebServerConfig(cfg *WebServerConfig){
	cfg.Port = 8015
	cfg.Address = "127.0.0.1"
	cfg.Config = ""
}

func loadConfig(cfg *WebServerConfig) {
	cfgPath := cfg.Config
	env := os.Getenv("EZGO_ENV")
	var err error
	if env == "" { env = "local" }
	if env == "heroku" { 
		cfg.Address = "0.0.0.0"
		cfg.Port, err = strconv.Atoi(os.Getenv("PORT"))
		fmt.Printf("HEROKU PORT and Address: %v %v\n", os.Getenv("PORT"), cfg.Address)
	}	
	if env != "heroku" && len(cfgPath) == 0 { cfgPath = "./.env." + env }	
	if env != "heroku" { err = godotenv.Load(cfgPath) }
	if err != nil {
		log.Fatalf("Error loading %s file", cfgPath)
	}
	fmt.Printf("%v %v %v\n", env, os.Getenv("EZGO_VERSION"), os.Getenv("EZGO_DB_URI"))
}

func ParseWebServerConfig(cfg *WebServerConfig) {
	// load default
	loadDefaultWebServerConfig(cfg)
	// load from parse
	flag.IntVar(&cfg.Port, "port", cfg.Port, "Web Server Port")
	flag.StringVar(&cfg.Address, "address", cfg.Address, "Web Server Address")
	flag.StringVar(&cfg.Config, "config", cfg.Config, "Config Path")
	flag.Parse()	
}

func startWebServer(cfg *WebServerConfig) {
	router := httprouter.New()
	addr := cfg.Address + ":" + strconv.Itoa(cfg.Port)
	router.GET("/", Index)
	server := http.Server{
		Addr: addr,
		Handler: router,
	}
	fmt.Printf("The web server is listen at http://%v\n", addr)
	log.Fatal(server.ListenAndServe())	
}

// Init : initialize the web server
func Init() {
	cfg := WebServerConfig{}
	ParseWebServerConfig(&cfg)	
	loadConfig(&cfg)
	startWebServer(&cfg)
}


// Helpers

func GenerateHTML(w *http.ResponseWriter, title string) {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(MainLayout())
	check(err)
	data := struct {
		Title string
	}{
		Title: title,
	}
	err = t.Execute(*w, data)	
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	GenerateHTML(&w, "Index")
}

func MainLayout() string{
	const tpl = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>{{.Title}}</title>
		</head>
		<body>
			<h1>{{.Title}}</h1>
		</body>
	</html>`
	return tpl
}

// type EmptyInterface interface {
// 	StructToMap(*map[string]interface{})
// }

// type Config struct {
// 	Port string
// 	Address string
// 	Config string
// }

// func (c Config) StructToMap(m *map[string]interface{}){
// 	ParseStructToMap(c, m)
// }


// type User struct {
// 	ID uint64
// 	FirstName string
// 	LastName string
// 	Login string
// 	PasswordHashed string
// }

// func (u User) StructToMap(m *map[string]interface{}) {
// 	ParseStructToMap(u, m)
// }

// func LoadConfig() {

// }

// func ParseConfig(cfg *Config){
// 	fmt.Print("parse config")
// 	mok := make(map[string]interface{})
// 	cfg.StructToMap(&mok)
// 	defaultConfig := make(map[string]interface{})
// 	defaultConfig["Port"] = 8081
// 	defaultConfig["Address"] = "127.0.0.1"
// }

// func SetStructField(f EmptyInterface, field string, value interface{}) {
// 	rf := reflect.TypeOf(f)
// 	fmt.Printf("Type: %v \n", rf)
// 	if rf.Kind() != reflect.Struct {
// 		panic("expects a struct")
// 	}

// 	reflect.ValueOf(&f).Elem().Set(reflect.ValueOf(value))
// }

// func ParseStructToMap(s EmptyInterface, m *map[string]interface{} ) {
// 	j, _ := json.Marshal(s)
// 	json.Unmarshal(j, m)
// }