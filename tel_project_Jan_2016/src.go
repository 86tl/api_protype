package main

import (
        "fmt"
        "net/http"
        "github.com/zenazn/goji"
        "github.com/zenazn/goji/web"
	"github.com/siddontang/go-mysql/client"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	conn, _ := client.Connect("127.0.0.1:3306", "root", "", "tel_project_feb_2016")
        //conn.Ping()
	res, _ := conn.Execute(`insert into tel_res (type) values (`+ string(c.URLParams["type"])+`)`)
	flag :=0
	if  res.InsertId >0{
		flag =1
	}
	//if err != nil{
	//	fmt.Println(err)
	//}
	fmt.Fprint(w, flag)
}

func main() {
        goji.Get("/save_type/:type", hello)
        goji.Serve()
}

//package main
//
//import (
//	//"fmt"
//	"github.com/go-martini/martini"
//        "github.com/siddontang/go-mysql/client"
//)
//func main() {
//  m := martini.Classic()
//  //m.Get("/", func() string {
//  //  return "Hello world!"
//  //})
//	m.Get("/save_type/:type", func(params martini.Params) string {
//	conn, _ := client.Connect("127.0.0.1:3306", "root", "", "tel_project_feb_2016")
//        //conn.Ping()
//	res, _ := conn.Execute(`insert into tel_res (type) values (`+ string(params["type"])+`)`)
//	//flag :=0
//	if  res.InsertId >0{
//		return "1"
//	    }else
//	{
//		return "0"
//	}
//		  //s := "'"+string(flag)+"'"
//		 //return "1111"
//  //return string(flag)
//      })
// m.RunOnAddr(":8000")
//}