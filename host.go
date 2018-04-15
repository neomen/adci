package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

var (
	path      = os.Getenv("NGINX_BASE")
	available = os.Getenv("NGINX_AVAILABLE")
	enabled   = os.Getenv("NGINX_ENABLED")
	fullPath  string
	aFile     string
	eFile     string
	output    string
)

// Nginx - Host structure
type Nginx struct {
	HostName   string
	Root       string
	DomainName string
}

// IndexFile - Host structure
type IndexFile struct {
	HostName string
}

//
func restartNginx() bool {
	fmt.Println("restartNginx")
	runcmd("service nginx restart")
	runcmd("service php7.1-fpm restart")
	return true
}

func createProjectPath(hostName string) {
	var path = os.Getenv("NGINX_BASE")
	fmt.Println("createProjectPath")
	var webPath = fmt.Sprintf("%s/%s/%s", path, hostName, "web")
	if _, err := os.Stat(webPath); os.IsNotExist(err) {
		os.MkdirAll(webPath, 0755)
	}
	var devPath = fmt.Sprintf("%s/%s/%s", path, hostName, "dev")
	if _, err := os.Stat(devPath); os.IsNotExist(err) {
		os.MkdirAll(devPath, 0755)
	}
	var statePath = fmt.Sprintf("%s/%s/%s", path, hostName, "stage")
	if _, err := os.Stat(statePath); os.IsNotExist(err) {
		os.MkdirAll(statePath, 0755)
	}
	var backupPath = fmt.Sprintf("%s/%s/%s", path, hostName, "backup")
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		os.MkdirAll(backupPath, 0755)
	}
	var logPath = fmt.Sprintf("%s/%s/%s", path, hostName, "log")
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		os.MkdirAll(logPath, 0755)
	}
}
func crateIndex(hostName string) {
	path = os.Getenv("NGINX_BASE")
	var tmplFile = "tmpl/dashboard.tmpl"
	t, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Print(err)
	}
	arrayEnv := [3]string{"web", "stage", "dev"}

	var indexFile = fmt.Sprintf("%s/%s/%s/index.html", path, hostName, "web")
	data := IndexFile{
		HostName: hostName,
	}
	for index, element := range arrayEnv {
		fmt.Println(index)
		fmt.Println(element)
		var indexFile = fmt.Sprintf("%s/%s/%s/index.html", path, hostName, element)

		f, err := os.Create(indexFile)
		if err != nil {
			log.Println("create file: ", err)
		}
		err = t.Execute(f, data)
		if err != nil {
			log.Print("execute: ", err)
		}
		f.Close()
	}

	f, err := os.Create(indexFile)
	if err != nil {
		log.Println("create file: ", err)
	}
	err = t.Execute(f, data)
	if err != nil {
		log.Print("execute: ", err)
	}
	defer f.Close()
}
func crateNginxConfig(hostName string, tmpl string) bool {
	fmt.Println("crateNginxConfig")
	var (
		path      = os.Getenv("NGINX_BASE")
		available = os.Getenv("NGINX_AVAILABLE")
		enabled   = os.Getenv("NGINX_ENABLED")
	)
	var tmplFile = "tmpl/" + os.Getenv("NGINX_DEFAULT_CONFIG")
	data := Nginx{
		HostName:   hostName,
		Root:       path + "/" + hostName,
		DomainName: os.Getenv("NGINX_BASE_DOMAIN"),
	}
	//@TODO add log rorater
	t, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Print(err)
		return false
	}
	var afile = available + "/" + hostName + ".conf"
	var efile = enabled + "/" + hostName + ".conf"
	f, err := os.Create(afile)
	if err != nil {
		log.Println("create file: ", err)
		return false
	}
	defer f.Close()
	runcmd("ln -s " + afile + " " + efile)
	err = t.Execute(f, data)
	if err != nil {
		log.Print("execute: ", err)
		return false
	}
	return true
}

func addHost(hostName string, tmpl string) bool {
	fmt.Println("addHost")
	// if tmpl == "d7" {
	// 	var tmplFile = "nginx-d7.tmpl"
	// }

	createProjectPath(hostName)
	crateNginxConfig(hostName, tmpl)
	crateIndex(hostName)
	//restartNginx()
	//tpl.Execute(f, data)
	//fmt.Printf("%s", output)

	return true
}
