package swagger

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func MessageReceive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")

	byt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("raw body", string(byt))

}
