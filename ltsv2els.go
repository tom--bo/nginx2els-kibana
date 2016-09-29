package main

import (
	"bufio"
	"net/http"
    "bytes"
	"os"
    "fmt"
    "strings"
//    "strconv"
    "encoding/json"
)

type NginxLog struct {
	Time         string `json:"time"`
	Host         string `json:"host"`
	Forwardedfor string `json:"forwardedfor"`
	Req          string `json:"req"`
	Method       string `json:"method"`
	Uri          string `json:"uri"`
	Status       string `json:"status"`
	Size         string `json:"size"`
	Referer      string `json:"referer"`
	UA           string `json:"ua"`
	Reqtime      string `json:"reqtime"`
	Runtime      string `json:"runtime"`
	Apptime      string `json:"apptime"`
	Cache        string `json:"cache"`
	Vhost        string `json:"vhost"`
}

func main() {
	var fp *os.File
	var err error
    //fmt.Println(os.Args[1])
    //fmt.Println(os.Args[2])

	if len(os.Args) < 2 {
		fp = os.Stdin

	} else {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)

		}
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		line := scanner.Text()
		items := strings.Split(line, "\t")
        logmap:=make(map[string]string)
        checkEmptyValue:=map[string]string{
            "time":"null",
            "host":"null",
            "forwardedfor":"null",
            "req":"null",
            "size":"null",
            "reqtime":"0",
            "runtime":"0",
            "apptime":"0",
        }
		for _, item := range items {
            val := strings.SplitN(item, ":", 2)
            if fill, ok:=checkEmptyValue[val[0]];ok{
                if val[1]=="-"{
                    val[1]=fill
                }
            }
            logmap[val[0]]=val[1]
		}
        log:=NginxLog{
            Time    : logmap["time"],
            Host    : logmap["host"],
            Forwardedfor : logmap["forwardedfor"],
            Req     : logmap["req"],
            Method  : logmap["method"],
            Uri     : logmap["uri"],
            Status  : logmap["status"],
            Size    : logmap["size"],
            Referer : logmap["referer"],
            UA      : logmap["ua"],
            Reqtime : logmap["reqtime"],
            Runtime : logmap["runtime"],
            Apptime : logmap["apptime"],
            Cache   : logmap["cache"],
            Vhost   : logmap["vhost"],
        }
		data, err := json.Marshal(log)
		if err != nil {
			fmt.Print("JSON marshaling failed: %s", err)
		}

        //fmt.Println(string(data))
        postJson(os.Args[2], data)

        // time.Sleep(10 * time.Millisecond)
		// b := new(bytes.Buffer)
		// json.NewEncoder(b).Encode(log)
		// _, err2 := http.Post(os.Args[2], "application/json; charset=utf-8", log)
		// r, err2 := http.Post(os.Args[2], "application/json; charset=utf-8",  bytes.NewBuffer(data))
		// if err2 != nil {
		// 	fmt.Println("Post failed: %s", err2)
		// }
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func postJson(url string, data []byte) {
        req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
        if err != nil {
            fmt.Println(err)
            panic(err)
        }
        // req.Header.Set("Content-Length", strconv.Itoa(len(data)))
        // req.Header.Set("Content-Length", strconv.Itoa(300))
        req.Header.Set("Content-Type", "application/json")

        client := http.Client{}
        resp, err2 := client.Do(req)
        if err2 != nil {
            fmt.Println(err2)
            panic(err2)
        }

        // fmt.Println(resp)
        // fmt.Println(resp.Body)
        defer resp.Body.Close()
}



