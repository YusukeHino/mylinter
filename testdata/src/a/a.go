package a

import (
	"log"
	"os"
)

func main() {
	fileNames := []string{"hoge.txt, foo.json"}
	for _, fn := range fileNames {
		f, err := os.Open(fn)
		if err != nil {
			log.Fatal("fail")
		}
		defer f.Close() // want "defer should not be used in for loop."
		f.Write([]byte{})
	}
}
