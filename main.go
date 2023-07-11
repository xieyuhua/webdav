package main

import (
	"flag"
	"log"
	"reflect"
	"os"
	"net/http"
	"strings"
    "fmt"
	"golang.org/x/net/webdav"
)

var bind *string
var path *string
var Prefix *string

func main() {
	bind = flag.String("p", "7777", "监听服务和端口")
    path = flag.String("d", "app,./", "目录 'app,/app;abc,/abc'")
    Prefix = flag.String("i", "/", "定义随机入口,保护入口安全, 例如：/abc/")
	flag.Parse()
	
	s_mux := http.NewServeMux();
	
    err := os.MkdirAll("workdir", os.ModePerm)
    if err != nil {
         fmt.Println("MkdirAll ", err)
        return
    }
	
	//清空目录
    err = os.RemoveAll("workdir")
    if err != nil {
        fmt.Println("RemoveAll", err)
        return
    }
        
	webdavs := make([]*webdav.Handler, 0)
	davpath := strings.Split(*path, ";")
	for _, dav := range davpath {
	    strArr := strings.Split(dav, ",")
    	Prefixname   := strArr[0]
    	pathDir      := strArr[1]
        webdavfs := &webdav.Handler{
        		Prefix:     *Prefix+Prefixname,
        		FileSystem: webdav.Dir(pathDir),
        		LockSystem: webdav.NewMemLS(),
        	}
		webdavs = append(webdavs, webdavfs)
		//创建目录
        err = os.MkdirAll("workdir/"+Prefixname, os.ModePerm)
        if err != nil {
             fmt.Println("dav MkdirAll ", err)
            return
        }
	}
		log.Println(davpath)
	// mkdir dav2 dav1  清空目录和重新创建目录
	f := &webdav.Handler{
		Prefix:     *Prefix,
		FileSystem: webdav.Dir("workdir"),
		LockSystem: webdav.NewMemLS(),
	}
	
	s_mux.HandleFunc(*Prefix, func(w http.ResponseWriter, req *http.Request) {
		// 获取用户名/密码
// 		DoMethod(req)
		log.Println("ip:",req.RemoteAddr," url: ",req.RequestURI)
// 		username, password, ok := req.BasicAuth()
// 		if !ok {
// 			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		//验证用户名/密码
// 		if username != "a" || password != "b" {
// 			http.Error(w, "WebDAV: need authorized!", http.StatusUnauthorized)
// 			return
// 		}
// 		switch req.Method {
// 		case "PUT", "DELETE", "PROPPATCH", "MKCOL", "COPY", "MOVE":
// 			http.Error(w, "WebDAV: Read Only!!!", http.StatusForbidden)
// 			return
// 		}

        for _, fs := range webdavs {
    		if strings.HasPrefix(req.RequestURI, fs.Prefix) {
    			fs.ServeHTTP(w, req)
    			return
    		}
        }
        
		//根目录文件
		if strings.HasPrefix(req.RequestURI, f.Prefix) {
			f.ServeHTTP(w,  req)
			return
		}
		// else
		w.WriteHeader(404)
	})

	dav_addr := fmt.Sprintf(":%v", *bind)
	log.Println("Dav Server run ", dav_addr, *Prefix)
	err = http.ListenAndServe(dav_addr, s_mux)
	if (err != nil) {
		fmt.Println("dav server run error:", err)
	}
}

// 通过接口来获取任意参数
func DoFiled(input interface{}) {
    getType := reflect.TypeOf(input) //先获取input的类型
    fmt.Println("Type is :", getType.Name()) // Person
    fmt.Println("Kind is : ", getType.Kind()) // struct
    getValue := reflect.ValueOf(input)
    fmt.Println("Fields is:", getValue) //{王富贵 20 男}
    // 获取方法字段
    // 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
    // 2. 再通过reflect.Type的Field获取其Field
    // 3. 最后通过Field的Interface()得到对应的value
        for i := 0; i < getType.NumField(); i++ {
            field := getType.Field(i)
            value := getValue.Field(i).Interface() //获取第i个值
            fmt.Printf("字段名称:%-20s, 字段类型:%-20s, 字段数值:%-20s \n", field.Name, field.Type, value)
        }
}


// 通过接口来获取任意参数
func DoMethod(input interface{}) {
    getType := reflect.TypeOf(input) //先获取input的类型
    fmt.Println("Type is :", getType.Name()) // Person
    fmt.Println("Kind is : ", getType.Kind()) // struct
    // 通过反射，操作方法
    // 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
    // 2. 再公国reflect.Type的Method获取其Method
        for i := 0; i < getType.NumMethod(); i++ {
            method := getType.Method(i)
            fmt.Printf("方法名称:%-20s, 方法类型:%-20s \n", method.Name, method.Type)
        }
}

