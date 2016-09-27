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
		var log NginxLog
		for _, item := range items {
            val := strings.SplitN(item, ":", 2)
            // fmt.Println(val[0] + ": " + val[1])
			switch {
			case val[0] == "time":
                if val[1] == "-" {
                    log.Time = "null"
                    continue
                }
				log.Time = val[1]
				continue
			case val[0] == "host":
                if val[1] == "-" {
                    log.Host = "null"
                    continue
                }
				log.Host = val[1]
				continue
			case val[0] == "forwardedfor":
                if val[1] == "-" {
                    log.Forwardedfor = "null"
                    continue
                }
				log.Forwardedfor = val[1]
				continue
			case val[0] == "req":
                if val[1] == "-" {
                    log.Req = "null"
                    continue
                }
				log.Req = val[1]
				continue
			case val[0] == "method":
				log.Method = val[1]
				continue
			case val[0] == "uri":
				log.Uri = val[1]
				continue
			case val[0] == "status":
				log.Status = val[1]
				continue
			case val[0] == "size":
                if val[1] == "-" {
                    log.Size = "null"
                    continue
                }
				log.Size = val[1]
				continue
			case val[0] == "referer":
				log.Referer = val[1]
				continue
			case val[0] == "ua":
				log.UA = val[1]
				continue
			case val[0] == "reqtime":
                if val[1] == "-" {
                    log.Reqtime = "0"
                    continue
                }
				log.Reqtime = val[1]
				continue
			case val[0] == "runtime":
                if val[1] == "-" {
                    log.Runtime = "0"
                    continue
                }
				log.Runtime = val[1]
				continue
			case val[0] == "apptime":
                if val[1] == "-" {
                    log.Apptime = "0"
                    continue
                }
				log.Apptime = val[1]
				continue
			case val[0] == "cache":
				log.Cache = val[1]
				continue
			case val[0] == "vhost":
				log.Vhost = val[1]
				continue
			}
		}
		data, err := json.Marshal(log)
		if err != nil {
			fmt.Print("JSON marshaling failed: %s", err)
		}

        // fmt.Println(log)
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



