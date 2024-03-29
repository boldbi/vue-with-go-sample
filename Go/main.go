package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"fmt"
)

var appconfig map[string]interface{}

type EmbedConfig struct {
	DashboardId    string `json:"DashboardId"`
	ServerUrl      string `json:"ServerUrl"`
	EmbedType      string `json:"EmbedType"`
	Environment    string `json:"Environment"`
	SiteIdentifier string `json:"SiteIdentifier"`
}

func main() {
	go func() {
		http.HandleFunc("/getdetails", getdetails)
		http.HandleFunc("/authorizationserver", authorizationServer)

		if err := http.ListenAndServe(":8086", nil); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Go server is running on port 8086")
	select {}
}

func getdetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fileData, err := ioutil.ReadFile("embedConfig.json")
	if err != nil {
		fmt.Println("embedCondfig.json file is missing")

	}
	err = json.Unmarshal(fileData, &appconfig)

	// Create a custom struct to hold the specific properties you want to return.
	clientEmbedConfigData := EmbedConfig{
		DashboardId:    appconfig["DashboardId"].(string),
		ServerUrl:      appconfig["ServerUrl"].(string),
		SiteIdentifier: appconfig["SiteIdentifier"].(string),
		EmbedType:      appconfig["EmbedType"].(string),
		Environment:    appconfig["Environment"].(string),
	}

	jsonResponse, err := json.Marshal(clientEmbedConfigData)
	w.Write(jsonResponse)
}

func authorizationServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	if len(body) > 0 {
		if queryString, err := unmarshal(string(body)); err != nil {
			fmt.Println("error converting", err)
		} else {
			userMail := appconfig["UserEmail"].(string)
			serverAPIUrl := queryString.(map[string]interface{})["dashboardServerApiUrl"].(string)
			embedQueryString := queryString.(map[string]interface{})["embedQuerString"].(string)
			embedQueryString += "&embed_user_email=" + userMail
			timeStamp := time.Now().Unix()
			embedQueryString += "&embed_server_timestamp=" + strconv.FormatInt(timeStamp, 10)
			signatureString, err := getSignatureUrl(embedQueryString)
			embedDetails := "/embed/authorize?" + embedQueryString + "&embed_signature=" + signatureString
			query := serverAPIUrl + embedDetails
			result, err := http.Get(query)
			if err != nil {
				fmt.Println(err)
			}
			response, err := ioutil.ReadAll(result.Body)
			if err != nil {
				log.Fatalln(err)
			}
			w.Write(response)
		}
	}
}

func getSignatureUrl(queryData string) (string, error) {
	embedSecret := appconfig["EmbedSecret"].(string)
	encoding := ([]byte(embedSecret))
	messageBytes := ([]byte(queryData))
	hmacsha1 := hmac.New(sha256.New, encoding)
	hmacsha1.Write(messageBytes)
	sha := base64.StdEncoding.EncodeToString(hmacsha1.Sum(nil))
	return sha, nil
}

func unmarshal(data string) (interface{}, error) {
	var iface interface{}
	decoder := json.NewDecoder(strings.NewReader(data))
	decoder.UseNumber()
	if err := decoder.Decode(&iface); err != nil {
		return nil, err
	}
	return iface, nil
}
