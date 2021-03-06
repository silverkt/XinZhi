package main

import (
	"net/http"
	"html/template"
    "io/ioutil"
	// "os"
	"time"

	"fmt"
)

func test1Handler(w http.ResponseWriter, r *http.Request) {
	t := template.New("第一个模板").Delims("[[", "]]") //创建一个模板,设置模板边界
	t, _ = t.Parse("hello,[[.UserName]]\n")       //解析模板文件
	data := map[string]interface{}{"UserName": template.HTML("<script>alert('you have been pwned')</script>")}
	// 此处template.HTML 是类型，加括号的意思是强制转换字符串伟 temlate.HTML类型，不转义
	//w.Header().Set("Content-Type","text/html");
	t.Execute(w, data) //执行模板的merger操作，并输出到控制台
	fmt.Println(t.Name(), "\n\n")
}

func test2Handler(w http.ResponseWriter, r *http.Request) {
	t2 := template.New("第二个模板")                         //创建模板
	t2.Funcs(map[string]interface{}{"tihuan": tihuan}) //向模板中注入函数
	bytes, _ := ioutil.ReadFile("tpl/test2.html")          //读文件
	template.Must(t2.Parse(string(bytes)))             //将字符串读作模板 并Must错误处理
	t2.Execute(w, map[string]interface{}{"UserName": "你好世界"})
	fmt.Println("\n", t2.Name(), "\n")

}

func test3Handler(w http.ResponseWriter, r *http.Request) {
	t3, _ := template.ParseFiles("tpl/test1.html") //将一个文件读作模板
	data := map[string]interface{}{"UserName": template.HTML("<script>alert('you have been pwned')</script>")}
	t3.Execute(w, data)
	fmt.Println("\n", t3.Name(), "\n") //模板名称
}

func test4Handler(w http.ResponseWriter, r *http.Request) {
	t4, _ := template.ParseGlob("tpl/test1.html") //将一个文件读作模板
	data := map[string]interface{}{"UserName": template.HTML("<script>alert('you have been pwned')</script>")}
	t4.Execute(w, data)
	fmt.Println(t4.Name())
}







func main() {
	http.HandleFunc("/test1", test1Handler);
	http.HandleFunc("/test2", test2Handler);
	http.HandleFunc("/test3", test3Handler);
	http.HandleFunc("/test4", test4Handler);

	http.ListenAndServe(":4000", nil);
}

//注入模板的函数
func tihuan(str string) string {
	return str + "-------" + time.Now().Format("2006-01-02");
}