package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

var (
	path      = "/var/www"
	available = "/etc/nginx/sites-available"
	enabled   = "/etc/nginx/sites-enabled"
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

//
func restartNginx() bool {
	runcmd("service nginx restart")
	runcmd("service php7.1-fpm restart")
	return true
}

func createProjectPath(hostName string) {
	var webPath = fmt.Sprintf("%s/%s/%s", path, hostName, "web")
	if _, err := os.Stat(webPath); os.IsNotExist(err) {
		os.MkdirAll(webPath, 0755)
	}
	var logPath = fmt.Sprintf("%s/%s/%s", path, hostName, "log")
	if _, err := os.Stat(webPath); os.IsNotExist(err) {
		os.MkdirAll(logPath, 0755)
	}
}

func addHost(hostName string, tmpl string) bool {
	// if tmpl == "d7" {
	// 	var tmplFile = "nginx-d7.tmpl"
	// }
	createProjectPath(hostName)
	var tmplFile = "nginx-d7.tmpl"
	data := Nginx{
		HostName:   hostName,
		Root:       path + "/" + hostName,
		DomainName: "clients.adciserver.com",
	}
	//@TODO add log rorater
	t, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Print(err)
		return false
	}

	f, err := os.Create("test.txt")
	if err != nil {
		log.Println("create file: ", err)
		return false
	}
	restartNginx()
	//tpl.Execute(f, data)
	//fmt.Printf("%s", output)
	err = t.Execute(f, data)
	if err != nil {
		log.Print("execute: ", err)
		return false
	}
	return true
}
