package generator

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)


func Generate(file *os.File, info []string) {
	var wg sync.WaitGroup
	writer := bufio.NewWriter(file)

	passwords := make(chan string, 6)

	go func() {
		for password := range passwords {
			_, err := writer.WriteString(password)
			if err != nil {
				log.Fatal(err)
			}

			writer.Flush()
		}
	}()

	wg.Add(len(info))

	for i := 0; i < len(info); i++ {
		go func(i int) {
			defer wg.Done()
			var b strings.Builder
			for j := 0; j < len(info); j++ {
				for k := 0; k < len(info); k++ {
					for l := 0; l < len(info); l++ {
						for m := 0; m < len(info); m++ {
							for n := 0; n < len(info); n++ {
								b.WriteString(fmt.Sprintf("%s%s%s%s%s%s\n",
									info[i], info[j], info[k], info[l], info[m], info[n]))
							}
						}
					}
				}
			}

			passwords <- b.String()
		}(i)
	}

	wg.Wait()
	close(passwords)

}
