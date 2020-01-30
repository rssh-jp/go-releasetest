# go-releasetest

# usage
```
package main

import (
	"log"

	"github.com/rssh-jp/go-releasetest"
)

func main() {
	r := releasetest.New(releasetest.OptionName("ore"), releasetest.OptionAge(55))
	log.Printf("%+v\n", r)
}
```
