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
)

 var appconfig map[string]interface{}
// func parseJSONFile(filePath string) error {
// 	// Read the JSON file
// 	fileData, err := ioutil.ReadFile(filePath)
	
// 	err = json.Unmarshal(fileData, &appconfig)
// 	if err != nil {
// 		return nil
// 	}
// 	log.Println(appconfig)
	
// 	if value, ok := appconfig["EmbedSecret"]; ok && value != nil {
// 		embedSecret = value.(string)
// 	}

// 	if value, ok := appconfig["UserEmail"]; ok && value != nil {
// 		userMail = value.(string)
// 	}

// 	return nil
// }

func main() {
	go func() {
		http.HandleFunc("/getdetails", getdetails)
		http.HandleFunc("/authorizationServer", authorizationServer)

		if err := http.ListenAndServe(":8086", nil); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Go server is running on port 8086")
	select {}
}

func getdetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fileData, err := ioutil.ReadFile("embedConfig.json")
	if err != nil {
		log.Println("embedConfig file is not found")

	}
	err = json.Unmarshal(fileData, &appconfig)
	jsonResponse, err := json.Marshal(appconfig)
	//log.Println(response)
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
			log.Println("error converting", err)
			http.Error(w, "Error converting JSON", http.StatusBadRequest)
			return
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
			//log.Println(query)
			result, err := http.Get(query)
			if err != nil {
				log.Println(err)
				http.Error(w, "Error calling API", http.StatusInternalServerError)
				return
			}
			//log.Println(result)
			response, err := ioutil.ReadAll(result.Body)
			if err != nil {
				log.Fatalln(err)
				http.Error(w, "Error reading response", http.StatusInternalServerError)
				return
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
