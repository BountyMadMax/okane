package lib

import (
	"fmt"
	"io"
	"os"
)

func ReadIcon(icon string) string {
	file, err := os.Open("node_modules/lucide-static/icons/" + icon + ".svg")

	if err != nil {
		fmt.Printf("Error reading icon %s : %s \n", icon, err)
		return ""
	}

	c := ""
	b := make([]byte, 8)

	for {
		_, err = file.Read(b)

		if err != nil && err != io.EOF {
			break
		}

		c += string(b)

		if err == io.EOF {
			break
		}
	}

	return c
}
