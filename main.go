package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/TwiN/go-color"
	QRCode "github.com/skip2/go-qrcode"
)

func LoadHomeHTMLFile(w http.ResponseWriter) (*template.Template, error) {
	view, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return nil, err
	}
	return view, nil
}

func GetHTMLData(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	var err error
	data := map[string]interface{}{"bytes": []byte{}, "message": ""}
	data["message"] = r.URL.Query().Get("message")

	if data["message"] != "" {
		data["bytes"], err = QRCode.Encode(data["message"].(string), QRCode.Medium, 256)
	}

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return nil, err
	}

	return data, nil
}

func Home(w http.ResponseWriter, r *http.Request) {
	template, err := LoadHomeHTMLFile(w)
	if err != nil {
		return
	}

	data, err := GetHTMLData(w, r)
	if err != nil {
		return
	}

	err = template.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
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
