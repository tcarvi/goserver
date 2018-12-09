/*
 *
 * Server for HTML, File and Text request
 *
 */
package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"log"
)

//  Implemenetation of http.Handler Interface: https://godoc.org/net/http#Handler
type THandler int
func (tHandler THandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

	// For debug: http.Request fields
	fmt.Println("Debug: r.URL.Path = " + r.URL.Path)
	fmt.Println("Debug: r.URL.Host = " + r.URL.Host)
	fmt.Println("Debug: r.Method = " + r.Method)
	fmt.Println("Debug: r.Proto = " + r.Proto)
	fmt.Println("Debug: r.Host = " + r.Host)
	
	if strings.HasPrefix(r.URL.Path, "/api"){
		fmt.Println("Debug: tServeAPI")
		tServeAPI(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".js"){
		fmt.Println("Debug: tServeJs")
		tServeJs(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/layout"){
		fmt.Println("Debug: tServeLayout")
		tServeLayout(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".css"){
		fmt.Println("Debug: tServeCss")
		tServeCss(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".svg"){
		fmt.Println("Debug: tServeSvg")
		tServeSvg(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".html"){
		fmt.Println("Debug: tServeHtml")
		tServeHtml(w, r)
	} else if strings.HasSuffix(r.URL.Path, "/"){
		fmt.Println("Debug: tServeHtml")
		tServeHtml(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".ico"){
		fmt.Println("Debug: tServeIco")
		tServeIco(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".jpeg"){
		fmt.Println("Debug: tServeJpeg")
		tServeJpeg(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".png"){
		fmt.Println("Debug: tServePng")
		tServePng(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".pdf"){
		fmt.Println("Debug: tServePdf")
		tServePdf(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".file"){
		fmt.Println("Debug: tServeFile")
		tServeFile(w, r)
	}
}

func tServeAPI(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Host", "tcarvi.com")
	// Message
	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	w.Header().Set("Content-Language", "pt-br")
	// TCARVI headers
	w.Header().Set("TCARVI-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "../../../../../webapp"+r.URL.Path)

	// Response Headers
	// HTTP/1.1 200 OK
	// Accept-Ranges: bytes
	// Cache-Control: no-cache
	// Content-Language: pt-br
	// Content-Type: text/javascript; charset=utf-8
	// Host: tcarvi.com
	// Last-Modified: Sun, 09 Dec 2018 21:08:40 GMT
	// Tcarvi-Code: 00000000001
	// Transfer-Encoding: gzip
	// Date: Sun, 09 Dec 2018 21:10:32 GMT
	// Transfer-Encoding: chunked
}

func tServeJs(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Host", "tcarvi.com")
	// Message
	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	w.Header().Set("Content-Language", "pt-br")
	// TCARVI headers
	w.Header().Set("TCARVI-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "../../../../../webapp/"+strings.TrimSuffix(r.Host, ":8080")+r.URL.Path)
}

func tServeCss(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Host", "tcarvi.com")
	// Message
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	// TCARVI headers
	w.Header().Set("TCARVI-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "../../../../../webapp/"+strings.TrimSuffix(r.Host, ":8080")+r.URL.Path)

	// Response Headers
	// HTTP/1.1 200 OK
	// Accept-Ranges: bytes
	// Cache-Control: no-cache
	// Content-Type: text/css; charset=utf-8
	// Host: tcarvi.com
	// Last-Modified: Sun, 09 Dec 2018 21:09:02 GMT
	// Tcarvi-Code: 00000000001
	// Transfer-Encoding: gzip
	// Date: Sun, 09 Dec 2018 21:10:32 GMT
	// Transfer-Encoding: chunked
}

func tServeSvg(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Host", "tcarvi.com")
	// Message
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	w.Header().Set("Content-Language", "pt-br")
	// TCARVI headers
	w.Header().Set("TCARVI-Code", "00000000001")
	//Serve File
	http.ServeFile(w, r, "../../../../../webapp/svg"+r.URL.Path)
}

func tServeHtml(w http.ResponseWriter, r *http.Request) {
	// Protocol
	w.Header().Set("HTTP-Version", "HTTP/2.0") //Chrome OK
	w.Header().Set("X-Firefox-Spdy", "h2")     //Chrome OK
	w.Header().Set("Protocol", "HTTP/2.0")     //Chrome OK
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip") //Chrome OK
	w.Header().Set("Cache-Control", "no-cache") //Chrome OK
	w.Header().Set("Host", "tcarvi.com")    //Chrome OK
	w.Header().Set("Date", time.Now().Format(time.UnixDate))
	// time.UnixDate: "Sat Mar  4 22:08:50 BRT 2017"
	// time.RFC3339:  "2017-03-04T22:05:21-03:00"
	// time.ANSIC     "Sat Mar  4 22:06:41 2017"
	// Message
	w.Header().Set("Content-Type", "text/html; charset=utf-8") //Chrome OK
	w.Header().Set("Content-Language", "pt-br")                //Chrome OK
	// TCARVI headers
	w.Header().Set("TCARVI-Code", "00000000001")
	//Serve File
	http.ServeFile(w, r, "../../../../../webapp/"+strings.TrimSuffix(r.Host, ":8080")+"/index.html")

	// Response Headers
	// HTTP/1.1 200 OK
	// Accept-Ranges: bytes
	// Cache-Control: no-cache
	// Content-Language: pt-br
	// Content-Type: text/html; charset=utf-8
	// Date: Sun Dec  9 19:10:32 -02 2018
	// Host: tcarvi.com
	// Http-Version: HTTP/2.0
	// Last-Modified: Sun, 09 Dec 2018 21:09:34 GMT
	// Protocol: HTTP/2.0
	// Tcarvi-Code: 00000000001
	// Transfer-Encoding: gzip
	// X-Firefox-Spdy: h2
	// Transfer-Encoding: chunked
}

func tServeLayout(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Host", "tcarvi.com")
	// Message
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	// TCARVI headers
	w.Header().Set("TCARVI-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "../../../../../webapp"+r.URL.Path)
}

func tServeIco(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Host", "tcarvi.com")
	// Message
	w.Header().Set("Content-Type", "image/x-icon")
	// TCARVI headers
	w.Header().Set("TCARVI-Code", "00000000001")
	http.ServeFile(w, r, "../../../../../webapp/"+strings.TrimSuffix(r.Host, ":8080")+r.URL.Path)
}

func tServeJpeg(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Host", "tcarvi.com")
	// Message
	w.Header().Set("Content-Type", "image/jpeg")
	// TCARVI headers
	w.Header().Set("TCARVI-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "../../../../../webapp/jpeg"+r.URL.Path)
}

func tServePng(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Host", "tcarvi.com")
	// Message
	w.Header().Set("Content-Type", "image/png")
	// TCARVI headers
	w.Header().Set("TCARVI-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "../../../../../webapp/png"+r.URL.Path)
}

func tServePdf(w http.ResponseWriter, r *http.Request) {
		// Connection
		w.Header().Set("Host", "tcarvi.com")
		// Message
		w.Header().Set("Content-Type", "/pdf")
		// TCARVI headers
		w.Header().Set("TCARVI-Code", "00000000001")
		//Serve Files
	http.ServeFile(w, r, "../../../../../webapp/pdf"+r.URL.Path)
}
func tServeFile(w http.ResponseWriter, r *http.Request) {
		// Connection
		w.Header().Set("Host", "tcarvi.com")
		// Message
		w.Header().Set("Content-Type", "charset=utf-8")
		// TCARVI headers
		w.Header().Set("TCARVI-Code", "00000000001")
	http.ServeFile(w, r, "../../../../../webapp"+r.URL.Path)
}

// TODO: func database()
// 	db, err := sql.Open("postgres", "user=postgres dbname=np sslmode=disable")
// 	dealError(err)
// 	defer db.Close()

func main() {
	var d THandler
	err := http.ListenAndServe(":8080", d)
	// TODO HTTPS ACCESS: log.Fatal(srv.ListenAndServeTLS(certFile, keyFile string))
	if err != nil {
		log.Fatal(err)
	}
}