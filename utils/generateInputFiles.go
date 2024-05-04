package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

func CreateJSONFile() {
	jsonFile, err := os.Create("jsonInput.json")
	if err != nil {
		fmt.Printf("json file creation fail: %s", err.Error())
	}

	defer jsonFile.Close()

	ports := []string{"5000:5000"}
	volumes := []string{"/usercode/:/code"}
	links := []string{"database:backenddb"}

	web := models.Web{
		Build:   ".",
		Ports:   ports,
		Volumes: volumes,
		Links:   links,
	}

	environment := []string{
		"MYSQL_ROOT_PASSWORD=root",
		"MYSQL_USER=testuser",
		"MYSQL_PASSWORD=admin123",
		"MYSQL_DATABASE=backend",
	}

	volumes = []string{"/usercode/db/init.sql:/docker-entrypoint-initdb.d/init.sql"}

	database := models.Database{
		Image:       "mysql/mysql-server:5.7",
		Environment: environment,
		Volumes:     volumes,
	}

	services := models.Services{
		Web:      web,
		Database: database,
	}

	dockerCompose := models.DockerCompose{
		Version:  "3",
		Services: services,
	}

	out, err := json.Marshal(&dockerCompose)
	if err != nil {
		fmt.Printf("json encoding fail: %s", err.Error())
	}

	_, err = jsonFile.Write(out)
	if err != nil {
		fmt.Printf("writing data fail: %s", err.Error())
	}
}

func CreateYAMLFile() {
	yamlFile, err := os.Create("yamlInput.yml")
	if err != nil {
		fmt.Printf("yaml file creation fail: %s", err.Error())
	}

	defer yamlFile.Close()

	ports := []string{"5000:5000"}
	volumes := []string{"/usercode/:/code"}
	links := []string{"database:backenddb"}

	web := models.Web{
		Build:   ".",
		Ports:   ports,
		Volumes: volumes,
		Links:   links,
	}

	environment := []string{
		"MYSQL_ROOT_PASSWORD=root",
		"MYSQL_USER=testuser",
		"MYSQL_PASSWORD=admin123",
		"MYSQL_DATABASE=backend",
	}

	volumes = []string{"/usercode/db/init.sql:/docker-entrypoint-initdb.d/init.sql"}

	database := models.Database{
		Image:       "mysql/mysql-server:5.7",
		Environment: environment,
		Volumes:     volumes,
	}

	services := models.Services{
		Web:      web,
		Database: database,
	}

	dockerCompose := models.DockerCompose{
		Version:  "3",
		Services: services,
	}

	out, err := yaml.Marshal(&dockerCompose)
	if err != nil {
		fmt.Printf("yaml encoding fail: %s", err.Error())
	}

	_, err = yamlFile.Write(out)
	if err != nil {
		fmt.Printf("writing data fail: %s", err.Error())
	}
}
