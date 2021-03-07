package implants

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"text/template"

	nodes "github.com/TerminalJockey/Kudzu/Nodes"
)

//ImplantOps holds values for implant options
type ImplantOps struct {
	ImplantType, FileName, Arch string
	Listener                    nodes.Listener
}

//GenerateImplant takes an ImplantOps struct and generates a new implant
func GenerateImplant(Ops ImplantOps) {
	switch Ops.ImplantType {
	case "cmd":
		scriptbytes, err := ioutil.ReadFile("Scripts/simpleagent_cmd.kzs")
		if err != nil {
			log.Println(err)
		}
		//setup template buffer
		tmplbuf := new(bytes.Buffer)
		tmpl, err := template.New("").Parse(string(scriptbytes))
		type Filler struct {
			Laddr string
		}
		filler := Filler{
			Laddr: Ops.Listener.Listener.Addr().String(),
		}
		tmpl.Execute(tmplbuf, filler)
		tmpname := nodes.GenUID() + ".go"
		tmpfile, err := os.Create("tmp/" + tmpname)
		_, err = tmpfile.Write(tmplbuf.Bytes())
		if err != nil {
			log.Println(err)
		}
		buildcmd := exec.Command("go", "build", "-o", "tmp/"+Ops.FileName, "tmp/"+tmpname)
		buildcmd.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64")
		err = buildcmd.Run()
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("generated implant! check tmp/%s\n", Ops.FileName)
		tmpfile.Close()
		err = os.Remove("tmp/" + tmpname)
		if err != nil {
			log.Println(err)
		}
	case "psh":
		scriptbytes, err := ioutil.ReadFile("Scripts/simpleagent_psh.kzs")
		if err != nil {
			log.Println(err)
		}
		//setup template buffer
		tmplbuf := new(bytes.Buffer)
		tmpl, err := template.New("").Parse(string(scriptbytes))
		if err != nil {
			log.Println(err)
		}
		type Filler struct {
			Laddr string
		}
		filler := Filler{
			Laddr: Ops.Listener.Listener.Addr().String(),
		}
		tmpl.Execute(tmplbuf, filler)
		tmpname := nodes.GenUID() + ".go"
		tmpfile, err := os.Create("tmp/" + tmpname)
		_, err = tmpfile.Write(tmplbuf.Bytes())
		if err != nil {
			log.Println(err)
		}
		buildcmd := exec.Command("go", "build", "-o", "tmp/"+Ops.FileName, "tmp/"+tmpname)
		buildcmd.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64")
		err = buildcmd.Run()
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("generated implant! check tmp/%s\n", Ops.FileName)
		tmpfile.Close()
		err = os.Remove("tmp/" + tmpname)
		if err != nil {
			log.Println(err)
		}
	case "kdzshell_win":
		scriptbytes, err := ioutil.ReadFile("Scripts/kdzshell_win.kzs")
		if err != nil {
			log.Println(err)
		}
		//setup template buffer
		tmplbuf := new(bytes.Buffer)
		tmpl, err := template.New("").Parse(string(scriptbytes))
		type Filler struct {
			Laddr string
		}
		filler := Filler{
			Laddr: Ops.Listener.Listener.Addr().String(),
		}
		tmpl.Execute(tmplbuf, filler)
		tmpname := nodes.GenUID() + ".go"
		tmpfile, err := os.Create("tmp/" + tmpname)
		_, err = tmpfile.Write(tmplbuf.Bytes())
		if err != nil {
			log.Println(err)
		}

		buildcmd := exec.Command("go", "build", "-o", "tmp/"+Ops.FileName, "tmp/"+tmpname)
		buildcmd.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64")
		err = buildcmd.Run()
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("generated implant! check tmp/%s\n", Ops.FileName)
		tmpfile.Close()
		err = os.Remove("tmp/" + tmpname)
		if err != nil {
			log.Println(err)
		}
	case "sh":
		scriptbytes, err := ioutil.ReadFile("Scripts/simpleagent_sh.kzs")
		if err != nil {
			log.Println(err)
		}
		//setup template buffer
		tmplbuf := new(bytes.Buffer)
		tmpl, err := template.New("").Parse(string(scriptbytes))
		if err != nil {
			log.Println(err)
		}
		type Filler struct {
			Laddr string
		}
		filler := Filler{
			Laddr: Ops.Listener.Listener.Addr().String(),
		}
		tmpl.Execute(tmplbuf, filler)
		tmpname := nodes.GenUID() + ".go"
		tmpfile, err := os.Create("tmp/" + tmpname)
		_, err = tmpfile.Write(tmplbuf.Bytes())
		if err != nil {
			log.Println(err)
		}
		buildcmd := exec.Command("go", "build", "-o", "tmp/"+Ops.FileName, "tmp/"+tmpname)
		buildcmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64")
		err = buildcmd.Run()
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("generated implant! check tmp/%s\n", Ops.FileName)
		tmpfile.Close()
		err = os.Remove("tmp/" + tmpname)
		if err != nil {
			log.Println(err)
		}
	default:
		fmt.Println("implant types: cmd, psh, kdzshell")
	}
}
