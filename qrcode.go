/*
@Time : 2018/6/13 下午3:11
@Author : tengjufeng
@File : frame
@Software: GoLand
*/

package main

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"net/http"
	"strings"
)

//解析url
func createUrl(form map[string][]string) string {
	var url string
	if len(form) < 2 {
		url = form["url"][0]
	} else {
		slice := make([]string, 0, 10)

		for k, v := range form {
			if k == "url" {

				slice = append(slice, v[0])
			}
		}
		for key, value := range form {
			if key != "url" {
				slice = append(slice, key+"="+value[0])
			}
		}
		url = strings.Join(slice, "&")
	}
	return url
}

//路由执行方法
func qrcodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	//w.Header().Set("content-type", "application/json")             //返回数据格式是json
	//w.Header().Set("content-type", "image/png") //返回数据格式是json
	r.ParseForm() //解析参数，默认是不会解析的
	urls := createUrl(r.Form)

	createQrcode(urls, w)

}

//生成二维码，并返回
func createQrcode(url string, w http.ResponseWriter) {
	qrCode, _ := qr.Encode(url, qr.L, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 225, 225)

	//file, _ := os.Create("qr2.png")
	//defer file.Close()

	png.Encode(w, qrCode)
}
func main() {

	http.HandleFunc("/", qrcodes)             //设置访问路由
	err := http.ListenAndServe(":21000", nil) //设置监听端口
	if err != nil {
		fmt.Println("server error:", err)
	}
	//fmt.Println("server start http://localhost:21000")
}
