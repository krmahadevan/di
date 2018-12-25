package di

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"go/build"
	"net/http"
)

type Config struct {
	Enabled      bool
	DatabasePath string
	Port         string
}

func NewConfig() *Config {
	return &Config{
		Enabled:      true,
		DatabasePath: build.Default.GOPATH + "/src/github.com/krmahadevan/di/example.db",
		Port:         "8000",
	}
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PersonRepository struct {
	database *sql.DB
}

func (repository *PersonRepository) FindAll() []*Person {
	rows, _ := repository.database.Query(`SELECT id, name, age FROM people;`)
	defer func() {
		_ = rows.Close()
	}()

	var people []*Person

	for rows.Next() {
		var (
			id   int
			name string
			age  int
		)

		_ = rows.Scan(&id, &name, &age)

		people = append(people, &Person{
			Id:   id,
			Name: name,
			Age:  age,
		})
	}

	return people
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{database: database}
}

type PersonService struct {
	config     *Config
	repository *PersonRepository
}

func (service *PersonService) FindAll() []*Person {
	if service.config.Enabled {
		return service.repository.FindAll()
	}

	return []*Person{}
}

func NewPersonService(config *Config, repository *PersonRepository) *PersonService {
	return &PersonService{config: config, repository: repository}
}

type Server struct {
	config        *Config
	personService *PersonService
}

func (server *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/people", server.findPeople)

	return mux
}

func (server *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + server.config.Port,
		Handler: server.Handler(),
	}

	_ = httpServer.ListenAndServe()
}

func (server *Server) findPeople(writer http.ResponseWriter, request *http.Request) {
	people := server.personService.FindAll()
	bytes, _ := json.Marshal(people)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write(bytes)
}

func NewServer(config *Config, personService *PersonService) *Server {
	return &Server{
		config:        config,
		personService: personService,
	}
}

func ConnectDatabase(config *Config) (*sql.DB, error) {
	return sql.Open("sqlite3", config.DatabasePath)
}




