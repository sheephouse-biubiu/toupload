package uploadservice

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method != http.MethodPost {
		w.Write([]byte("限制post请求"))
		return
	}
	// w.Write([]byte("upload create"))
	fmt.Print(r.Header.Get("content-type"))
	if r.Header.Get("content-type") != "application/json" {
		w.Write([]byte("请求body数据需要格式:json"))
		return
	}
	var c CreateRequest
	if r.Body == nil {
		return
	}
	fmt.Println(r.Body)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&c)
	//json.Unmarshal(r.Body, &c)
	fmt.Println(c.FileName)
	fmt.Println(c.Slices)
	fmt.Println(c.Total)
	w.Write([]byte(c.FileName))
}

func Upload(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("upload...."))
}

func Merge(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("merge func..."))
}

type RouteHandler func(http.ResponseWriter, *http.Request)
type RouteInfo struct {
	path    string
	handler RouteHandler
}

var RouteList []RouteInfo

func init() {
	RouteList = append(RouteList, RouteInfo{"/easyUpload/create", Create})
	RouteList = append(RouteList, RouteInfo{"/easyUpload/upload", Upload})
	RouteList = append(RouteList, RouteInfo{"/easyUpload/merge", Merge})
}
