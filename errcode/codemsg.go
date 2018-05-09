//
// +build ignore

package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/alecthomas/template"
)

const codeMsgTmpl = `package errcode

var msg = map[int]string{
{{range $key, $value := .}}    {{$key}}: ` + "`{{$value}}`," + `
{{end}}}
`

func main() {
	b, err := ioutil.ReadFile("err_code_msg.txt")
	if err != nil {
		log.Fatalln(err)
	}
	replacer := strings.NewReplacer("\r\n", "\n", "\r", "\n", "\t", " ")
	s := replacer.Replace(string(b))
	lines := strings.Split(s, "\n")
	msg := make(map[int]string)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		ss := strings.SplitN(line, " ", 2)
		if len(ss) != 2 {
			continue
		}
		key, err := strconv.Atoi(ss[0])
		if err != nil {
			continue
		}
		msg[key] = ss[1]
	}
	var buf bytes.Buffer
	t := template.Must(template.New("errcodemsg").Parse(codeMsgTmpl))
	if err = t.Execute(&buf, msg); err != nil {
		log.Fatalln(err)
	}
	if err = ioutil.WriteFile("err_code_msg.go", buf.Bytes(), 0644); err != nil {
		log.Fatalln(err)
	}
	log.Printf("done!")
}
