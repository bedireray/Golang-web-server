package main

import (
	"net/http"
	"html/template"
)

type mainpage struct {
	Title string
	News string
}


func mainpagehandler(w http.ResponseWriter,r *http.Request)  {
	p := mainpage{Title:"deneme", News:"hebe hehbe"}
	t,_:=template.ParseFiles("mainpage.html")
	t.Execute(w,p)
	
}
func main()  {
	http.HandleFunc("/",mainpagehandler)
	http.ListenAndServe(":80",nil)
}

