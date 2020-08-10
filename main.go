package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

// HandleReCAPTCHA HandleReCAPTCHA
func HandleReCAPTCHA(router *mux.Router) {
	router.HandleFunc("/recaptcha", func(w http.ResponseWriter, r *http.Request) {
		vf := r.URL.Query().Get("vf")

		v := url.Values{}
		v.Set("secret", "6LebxbwZAAAAACiU0uWV-MqTrwJ3OcGcO0019TVh")
		v.Set("response", vf)
		resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", v)
		if err != nil {
			fmt.Println(err)
			return
		}

		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
	})
	router.HandleFunc("/reclient/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/tmp/recaptcha.html")
	})
}

func main() {
	router := mux.NewRouter()

	HandleReCAPTCHA(router)

	fmt.Println("listen at :8888")
	http.ListenAndServe(":8888", router)
}
