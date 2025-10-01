package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"log-crawler/internal/common"
)

func main() {
	println(">>>>> Starting log crawler ==================================================")
	fmt.Println("")
	fmt.Printf(">>>>> Starting log crawler ==================================================\n")
	fmt.Printf("%s\n", time.Now().Format("2006-01-02 15:04:05.000"))
	fmt.Println("")

	if len(os.Args) < 1 {
		println("Param is missing")
		os.Exit(2)
	}

	moduleName := os.Args[1]

	// Read configuration from JSON file
	config, err := readConfig("data/" + moduleName + ".json")
	if err != nil {
		println("Error reading config:", err.Error())
		return
	}

	// Process each service in the configuration
	for i, service := range config.Services {
		if !service.Enable {
			println(" >>>>> Processing service " + fmt.Sprintf("%d", i+1) + ": skipped")
			continue
		}
		println(" >>>>> Processing service", i+1)
		processService(config.URLPrefix, config.Token, service)
		println(" <<<<< End service", i+1)
		println("")
		time.Sleep(1 * time.Second)
	}
}

// readConfig reads the service-list.json file and returns the configuration
func readConfig(filename string) (*common.Config, error) {
	println("Reading file:", filename)

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config common.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &config, nil
}

// processService handles a single service request
func processService(urlPrefix, token string, service common.Service) {
	url := urlPrefix + service.Path
	payload := strings.NewReader(service.Param)

	req, err := http.NewRequest(service.Method, url, payload)

	if err != nil {
		println("error:", err.Error())
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	client := &http.Client{}

	fmt.Println(url)
	println("Sending request to:", url)
	res, startTime, endTime, err := sendRequest(client, req)
	if err != nil {
		println("error:", err.Error())
		return
	}
	defer res.Body.Close()

	println("Response status:", res.Status)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		println("error:", err.Error())
		return
	}

	if res.StatusCode != 200 {
		fmt.Println("client |  | [] INFO | Client Request | start |" + startTime.Format("2006-01-02|15:04:05.000"))
		fmt.Println("client |  | [] INFO | Client Request | end |" + endTime.Format("2006-01-02|15:04:05.000"))
		fmt.Println(res.Header.Get("X-Performance-Tuning-StartLog"))
		fmt.Println(res.Header.Get("X-Performance-Tuning-EndLog"))
		fmt.Println("Error ", res.StatusCode)
		fmt.Println("body: ", string(body))
		fmt.Println("")

		return
	}

	prxStartArr := strings.Split(res.Header.Get("X-Performance-Tuning-StartLog"), "|")
	prxEndArr := strings.Split(res.Header.Get("X-Performance-Tuning-EndLog"), "|")

	proxyStartLog := prxStartArr[0] + "|" + prxStartArr[1] + "| [] INFO |" + strings.Split(prxStartArr[2], "- start")[0] + "| start |" + prxStartArr[3] + "|" + prxStartArr[4]
	proxyEndLog := prxEndArr[0] + "|" + prxEndArr[1] + "| [] INFO |" + strings.Split(prxEndArr[2], "- end")[0] + "| end |" + prxEndArr[3] + "|" + prxEndArr[4]

	id := strings.TrimSpace(prxStartArr[1])

	clientLogStart := "client | " + id + " | [] INFO | Client Request | start |" + startTime.Format("2006-01-02|15:04:05.000")
	clientLogEnd := "client | " + id + " | [] INFO | Client Request | end |" + endTime.Format("2006-01-02|15:04:05.000")

	fmt.Println(clientLogStart)
	fmt.Println(clientLogEnd)
	fmt.Println(proxyStartLog)
	fmt.Println(proxyEndLog)

	var printBody string

	if len(body) > 0 {
		printBody = "{}"
	} else {
		printBody = "null"
	}

	println("body: ", printBody)

	if id != "" {
		fmt.Println(getLog(urlPrefix, id, token))
	}
}

func sendRequest(client *http.Client, req *http.Request) (*http.Response, time.Time, time.Time, error) {
	start := time.Now()

	res, err := client.Do(req)

	end := time.Now()

	fmt.Printf("%d | client time: %.3f s\n", res.StatusCode, end.Sub(start).Seconds())

	if err != nil {
		return nil, start, end, err
	}

	return res, start, end, nil
}

func getLog(urlPrefix string, id string, token string) string {
	url := urlPrefix + "/log/get?id=" + id
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		println("getLog error 1")
		return err.Error()
	}

	req.Header.Add("Authorization", token)

	println("getLog with id:", id)
	res, err := client.Do(req)
	println("getLog end", res.StatusCode)

	if err != nil {
		println("getLog error 2")
		return err.Error()
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		println("getLog error 3")
		return err.Error()
	}

	return string(body)
}
