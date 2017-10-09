package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/jniltinho/easyssh"
)

const (
	VERSION    = "0.4"
	BUILD_DATE = "20160113"
	SOURCE     = "https://gist.github.com/jniltinho/347a1c32f9781808c81b"
)

/*
### Build (Linux, Windows, Mac)
go get github.com/jniltinho/easyssh
GOOS=linux GOARCH=amd64 go build -v mikrotik_backup.go
GOOS=windows GOARCH=386 go build -v -o mikrotik_backup.exe mikrotik_backup.go
GOOS=darwin GOARCH=amd64 go build -v mikrotik_backup.go

## How to use:
./mikrotik_backup --bk --ip=192.168.X.X --user=admin --pass=password
*/

var (
	mk_bk     bool
	mk_ip     string
	mk_pass   string
	mk_port   string = "22"
	mk_user   string = "admin"
	mk_folder string = "."
	mk_cmd    string = "export compact"
	mk_time   string = time.Now().Format("20060102_1504")
)

func mk_backup() {

	file_name := fmt.Sprintf("rb_%s-%s.rsc", mk_ip, mk_time)
	ssh := &easyssh.MakeConfig{User: mk_user, Server: mk_ip, Password: mk_pass, Port: mk_port}

	response, err := ssh.Run(mk_cmd)
	// Handle errors
	if err != nil {
		log.Println("Can't run remote command: ", err.Error())
	} else {
		//print(response)
		filename := fmt.Sprintf("%s/%s", mk_folder, file_name)
		println(response)
		log.Println(fmt.Sprintf("Run Backup RB %s FILE: %s.gz", mk_ip, filename))

		var buf bytes.Buffer
		gz := gzip.NewWriter(&buf)
		gz.Write([]byte(response))
		gz.Close()
		ioutil.WriteFile(filename+".gz", buf.Bytes(), 0666)
	}

}

func init() {

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Version:", VERSION)
		fmt.Fprintln(os.Stderr, "Built:", BUILD_DATE)
		fmt.Fprintln(os.Stderr, "Fonte:", SOURCE)
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "--bk")
		flag.PrintDefaults()
	}
	flag.StringVar(&mk_ip, "ip", "", "Mikrotik IP")
	flag.StringVar(&mk_port, "port", mk_port, "Mikrotik PORT")
	flag.StringVar(&mk_user, "user", mk_user, "Mikrotik USER")
	flag.StringVar(&mk_pass, "pass", "", "Mikrotik PASSWORD")
	flag.StringVar(&mk_folder, "folder", mk_folder, "Mikrotik FOLDER")
	flag.BoolVar(&mk_bk, "bk", false, "Backup Mikrotik")

}

func main() {

	flag.Parse()

	if mk_bk && mk_ip != "" && mk_pass != "" {
		mk_backup()
	}

}
