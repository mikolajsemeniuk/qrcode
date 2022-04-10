package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/TwiN/go-color"
	"github.com/skip2/go-qrcode"
)

func LoadHomeHTMLFile(_ http.ResponseWriter) (*template.Template, error) {
	view, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		return nil, err
	}
	return view, nil
}

func GetHTMLData(_ http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	var err error
	data := map[string]interface{}{"bytes": []byte{}, "message": ""}
	data["message"] = r.URL.Query().Get("message")

	if data["message"] != "" {
		data["bytes"], err = qrcode.Encode(data["message"].(string), qrcode.Medium, 256)
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}

func Home(w http.ResponseWriter, r *http.Request) {
	view, err := LoadHomeHTMLFile(w)
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}

	data, err := GetHTMLData(w, r)
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}

	err = view.Execute(w, data)
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}
}

func main() {
	http.HandleFunc("/", Home)

	fmt.Println(color.Ize(color.Green, "\nserver is running on port 8082...\n"))

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println("error: ", err)
	}
}
