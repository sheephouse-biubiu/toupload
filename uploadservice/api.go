package uploadservice

import (
	"fmt"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method != http.MethodPost {
		w.Write([]byte("限制post请求"))
		return
	}
	w.Write([]byte("upload create"))
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
