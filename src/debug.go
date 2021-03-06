package main

import "os"
import "fmt"
import "os/exec"
import "strings"
//import "github.com/fatih/color"
import "gopkg.in/cheggaaa/pb.v1"


func progress() {
	if peid.verbose == false {
		progressBar.Increment()
	}
}

func CreateProgressBar() {

	var full int = 46

	if peid.verbose == false {
		if peid.staged == true {
			full -= 10
		}

		progressBar = pb.New(full)
		progressBar.SetWidth(80)
		progressBar.Start()
	}
}

func verbose(str string, status string) {

	if peid.verbose == true {
		if status == "*" {
			BoldYellow.Print("[*] ")
			white.Println(str)	
		}else if status == "+" {
			BoldGreen.Print("[+] ")
			white.Println(str)	
		}else if status == "-" {
			BoldRed.Print("[-] ")
			white.Println(str)
		}else if status == "!"{
			BoldRed.Print("[!] ")
			white.Println(str)
		}else if status == "Y" {
			BoldYellow.Println(str)
		}else if status == "B" {
			BoldBlue.Println(str)
		}
	}
}

func _verbose(str string, value int32) {

	if peid.verbose == true {
		BoldYellow.Print("[*] ")
		white.Printf(str+" 0x%X\n",value)
	}
}


func CheckRequirements() {

	verbose("Checking requirments...","*")

	CheckMingw, mingwErr := exec.Command("sh", "-c", "i686-w64-mingw32-g++-win32 --version").Output()
	if !strings.Contains(string(CheckMingw), "Copyright") {
		BoldRed.Println()
		red.Println(string(CheckMingw))
		red.Println(mingwErr)
		os.Exit(1)
	}
	progress()
	CheckNasm, _ := exec.Command("sh", "-c", "nasm -h").Output()
	if !strings.Contains(string(CheckNasm), "usage:") {
		BoldRed.Println("\n\n[!] ERROR: nasm is not installed.")
		red.Println(string(CheckNasm))
		os.Exit(1)
	}
	progress()
	CheckStrip, _ := exec.Command("sh", "-c", "strip -V").Output()
	if !strings.Contains(string(CheckStrip), "Copyright") {
		BoldRed.Println("\n\n[!] ERROR: strip is not installed.")
		red.Println(string(CheckStrip))
		os.Exit(1)
	}
	progress()
	CheckXXD, _ := exec.Command("sh", "-c", "echo Amber|xxd").Output()
	if !strings.Contains(string(CheckXXD), "Amber") {
		BoldRed.Println("\n\n[!] ERROR: xxd is not installed.")
		red.Println(string(CheckMingw))
		os.Exit(1)
	}
	progress()
	CheckMultiLib, _ := exec.Command("sh", "-c", "apt-cache policy gcc-multilib").Output()
	if strings.Contains(string(CheckMultiLib), "(none)") {
		BoldRed.Println("\n\n[!] ERROR: gcc-multilib is not installed.")
		red.Println(string(CheckMultiLib))
		os.Exit(1)
	}
	progress()
	CheckMultiLibPlus, _ := exec.Command("sh", "-c", "apt-cache policy g++-multilib").Output()
	if strings.Contains(string(CheckMultiLibPlus), "(none)") {
		BoldRed.Println("\n\n[!] ERROR: g++-multilib is not installed.")
		red.Println(string(CheckMultiLibPlus))
		os.Exit(1)
	}
	progress()

}


func ParseError(Err error,ErrStatus string,Msg string){

	if Err != nil {
		
		//progressBar.Finish()
		fmt.Println("\n")
		BoldRed.Println(ErrStatus)
		fmt.Println(Err)
		if len(Msg) > 0 {
			BoldRed.Println(Msg)
		}
		clean()
		fmt.Println("\n")
		os.Exit(1)
	}
}


func clean() {

	exec.Command("sh", "-c", "rm core/ASLR/Mem.map").Run()
	progress()
	exec.Command("sh", "-c", "rm core/ASLR/iat/Mem.map").Run()
	progress()
	exec.Command("sh", "-c", "rm core/Fixed/Mem.map").Run()
	progress()
	exec.Command("sh", "-c", "rm core/Fixed/iat/Mem.map").Run()
	progress()
	exec.Command("sh", "-c", "rm stub.o").Run()
	progress()
	exec.Command("sh", "-c", "rm Payload").Run()
	progress()
	exec.Command("sh", "-c", "rm Payload.key").Run()
	progress()
	exec.Command("sh", "-c", "echo   > stub/payload.h").Run()
	progress()
	exec.Command("sh", "-c", "echo   > stub/key.h").Run()
	progress()

}
