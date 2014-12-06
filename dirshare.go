package main

import (
		"fmt"
		"net/http"
		"os"
	)


var config struct {
	default_port string
}


func init() {
	config.default_port = ":8080"
//	ConsoleArguments := os.Args[1:]
}

func checkdirexists(path string) (bool, error) {
	fi, err := os.Stat(path)
	
	if err != nil { return false, err }
		
	if(fi.IsDir()) {
		if err == nil { return true, nil }
		if os.IsNotExist(err) { return false, nil }
		return false, err
	} else {
		return false, err		
	}
	return false, err
}





func main() {
	ConsoleArguments := os.Args[1:]

	if(len(ConsoleArguments)==0) {
		fmt.Println("usage: dirshare <dirpath>")
	} else {
		DirectoryShare := ConsoleArguments[0]
		// check if the directory exists
		direxists, error := checkdirexists(DirectoryShare)
		if(direxists) {
			fmt.Println("Sharing directory: ", DirectoryShare , " via port", config.default_port) 
			http.ListenAndServe(config.default_port, http.FileServer(http.Dir(DirectoryShare)))
		} else {
			if(error != nil) {
				fmt.Println("Facing problems while accessing directory: ", DirectoryShare , " error: ", error)
			} else {
				fmt.Println("Directory '",DirectoryShare ,"' not found")
			}
		}

	}
}
