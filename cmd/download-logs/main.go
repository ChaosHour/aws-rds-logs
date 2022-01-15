package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

var (
	dbRegion = flag.String("r", "", "--region Example: us-west-1")
	dbIdent  = flag.String("d", "", "db-instance-identifier Example: aurora-test-copy-stg-0")
	logType  = flag.String("l", "", "log-file-name Example: slowquery/mysql-slowquery.log.2021-09-08.20")
)

func main() {
	flag.Parse()

	args := []string{
		"aws",
		"rds",
		"--region", *dbRegion,
		"download-db-log-file-portion",
		"--db-instance-identifier", *dbIdent,
		"--log-file-name", *logType,
		"--output", "text",
	}

	cmd := exec.Command(args[0], args[1:]...)
	//fmt.Println("executing:", (args))

	fmt.Printf("Executing command:\n%v\n\n", args)

	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Something went wrong, please review and try again.: %v", err)
	}
	fmt.Printf("%s\n", b)
	//flag.PrintDefaults()
}
