package main

import (
	"fmt"
	"os"
	"os/exec"
	"errors"
)

func errHandler(err error) {
	if err != nil { panic(err) }
}


func generateSSH() {
	folder := "./keys";
	fmt.Print(folder);
}


func sign(file string) string {
	cmd := exec.Command( fmt.Sprintf( "cat %s | ssh-keygen -Y sign -n file -f id_rsa > %s.sig", file, file ) );
	stdout, _ := cmd.Output();

	return string(stdout);
}

func validate(file string) string {
	cmd := exec.Command( fmt.Sprintf( "cat %s | ssh-keygen -Y check-novalidate -f id_rsa.pub -n file -s %s.sig", file, file ) );
	stdout, _ := cmd.Output();

	return string(stdout);
}


func main() {
	args := os.Args[1:];
	
	flag := args[0];
	path := args[1];

	generateSSH();

	if flag == "-f" || flag == "--file" {		// If flag is file
		fileInfo, err := os.Stat(path);
		errHandler(err);

		if fileInfo.IsDir() {
			errHandler( errors.New("flag set to file, but file is not probided") );
		}

		fmt.Println(sign(path));
		fmt.Println(validate(path));
	}

	



}