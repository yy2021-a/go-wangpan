package handler

import (
	"io"
	"net/http"
	"os"
)

//处理文件上传：uploadhandler
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传html页面
		data, err := os.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internet server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//接受文件流及存储到本地文件

	}
}
