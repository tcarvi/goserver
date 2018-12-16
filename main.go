/*
 * Server for HTML, File, Texts and Database requests
 */
package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"log"
)

type THandler int
func (tHandler THandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

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
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "github.com/tcarvi"+r.URL.Path)
}

func tServeJs(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Host", "tcarvi.com")
	// Message
	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	w.Header().Set("Content-Language", "pt-br")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "github.com/tcarvi/"+strings.TrimSuffix(r.Host, ":8080")+r.URL.Path)
}

func tServeCss(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Host", "tcarvi.com")
	// Message
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "github.com/tcarvi/"+strings.TrimSuffix(r.Host, ":8080")+r.URL.Path)
}

func tServeSvg(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Host", "tcarvi.com")
	// Message
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	w.Header().Set("Content-Language", "pt-br")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve File
	http.ServeFile(w, r, "github.com/tcarvi/svg"+r.URL.Path)
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
	// Message
	w.Header().Set("Content-Type", "text/html; charset=utf-8") //Chrome OK
	w.Header().Set("Content-Language", "pt-br")                //Chrome OK
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve File
	http.ServeFile(w, r, "github.com/tcarvi/"+strings.TrimSuffix(r.Host, ":8080")+"/index.html")
}

func tServeLayout(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Host", "tcarvi.com")
	// Message
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "github.com/tcarvi"+r.URL.Path)
}

func tServeIco(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Host", "tcarvi.com")
	// Message
	w.Header().Set("Content-Type", "image/x-icon")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	http.ServeFile(w, r, "github.com/tcarvi"+strings.TrimSuffix(r.Host, ":8080")+r.URL.Path)
}

func tServeJpeg(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Host", "tcarvi.com")
	// Message
	w.Header().Set("Content-Type", "image/jpeg")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "github.com/tcarvi/jpeg"+r.URL.Path)
}

func tServePng(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Host", "tcarvi.com")
	// Message
	w.Header().Set("Content-Type", "image/png")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "github.com/tcarvi/png"+r.URL.Path)
}

func tServePdf(w http.ResponseWriter, r *http.Request) {
		// Connection
		w.Header().Set("Host", "tcarvi.com")
		// Message
		w.Header().Set("Content-Type", "/pdf")
		// To be handled by webapp
		w.Header().Set("key-Code", "00000000001")
		//Serve Files
	http.ServeFile(w, r, "github.com/tcarvi/pdf"+r.URL.Path)
}
func tServeFile(w http.ResponseWriter, r *http.Request) {
		// Connection
		w.Header().Set("Host", "tcarvi.com")
		// Message
		w.Header().Set("Content-Type", "charset=utf-8")
		// To be handled by webapp
		w.Header().Set("key-Code", "00000000001")
	http.ServeFile(w, r, "github.com/tcarvi"+r.URL.Path)
}

// TODO - Remove text "/tcarvi" from this code.

// TODO: func database() string
// 	db, err := sql.Open("postgres", "user=postgres dbname=np sslmode=disable")

// TODO: func answer(q Question) Answer
// a, err := ai.Answers(q)

func main() {
	var d THandler
	err := http.ListenAndServe(":8080", d)
	// TODO HTTPS ACCESS: log.Fatal(srv.ListenAndServeTLS(certFile, keyFile string))
	if err != nil {
		log.Fatal(err)
	}
}