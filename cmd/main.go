package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

var (
	signFile string
	signFilePub string
)

func init() {
	hdir, _ := os.UserHomeDir();
	signFile = path.Join(
		hdir, ".ssh", "ds_sign",
	);
	signFile = strings.ReplaceAll(signFile, "\\", "/");
	signFilePub = signFile + ".pub";
}

func errHandler(err error) {
	if err != nil { panic(err) }
}

func runCommand(c string, a ...string) string {
	cmd := exec.Command( fmt.Sprintf( c, a ) );
	
	stdout, err := cmd.Output();
	errHandler(err)
	return string(stdout);
}


func generateSSH() {
	// ext := path.Ext(signFile);
	// fmt.Println(signFile, "\t\t->\t\t", ext);

	if _, err := os.Stat(signFile); errors.Is(err, os.ErrNotExist) {
		fmt.Println("Creating new ssh-key file");

		fmt.Println(runCommand("ssh-keygen -f \"%s\" -N \"''\"", signFile));
	}
	fmt.Println("File is created or already exist");
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