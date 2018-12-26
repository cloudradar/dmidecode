// +build !windows

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/troian/dmidecode"
)

func main() {
	cmd := exec.Command("dmidecode")

	buf := bytes.Buffer{}
	cmd.Stdout = bufio.NewWriter(&buf)
	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}

	dmi, err := dmidecode.Unmarshal(bufio.NewReader(&buf))
	if err != nil {
		fmt.Println(err.Error())
	}

	raw, _ := dmi.Raw()

	data, err := json.MarshalIndent(raw, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(data))
	}
}
