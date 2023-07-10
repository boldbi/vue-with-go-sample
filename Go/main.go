package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	//"runtime/debug"
	"strconv"
	"strings"
	"time"
)

//Set EmbedSecret key from Bold BI Server. Please refer this link(https://help.syncfusion.com/bold-bi/on-premise/site-settings/embed-settings)
var embedSecret = "iUgYqIhFooyeI5HPRL9siUclfIAZy4eT"

//Enter your BoldBI Server credentials
var userMail = "dharshini.v@syncfusion.com"

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/parseJson", parseJSONHandler)
	http.HandleFunc("/getDetails", getEmbedDetails)

	go func() {
		if err := http.ListenAndServe(":8086", nil); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Go server is running on port 8086")

	// Keep the main goroutine alive
	select {}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("App is running!"))
}

// func parseJSONHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Context-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Println("Error reading request body:", err)
// 		http.Error(w, "Error reading request body", http.StatusBadRequest)
// 		return
// 	}

// 	var jsonData interface{}
// 	err = json.Unmarshal(body, &jsonData)
// 	if err != nil {
// 		log.Println("Error parsing JSON:", err)
// 		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
// 		return
// 	}

// 	// Process the received JSON data as needed
// 	// ...

// 	response := map[string]interface{}{
// 		"status":  "success",
// 		"message": "JSON data received and processed successfully",
// 	}

// 	jsonResponse, err := json.Marshal(response)
// 	if err != nil {
// 		log.Println("Error creating JSON response:", err)
// 		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Write(jsonResponse)
// }
func parseJSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Read the embedConfig.json file
	filePath := "embedConfig.json"
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("embedConfig file is not found")
		//http.Error(w, "Error reading", http.StatusInternalServerError)
		return
	}

	var jsonData interface{}
	err = json.Unmarshal(fileData, &jsonData)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		http.Error(w, "Error parsing JSON", http.StatusInternalServerError)
		return
	}

	if jsonData == nil {
		log.Println("Received JSON data is nil")
		http.Error(w, "Received JSON data is nil", http.StatusBadRequest)
		return
	}

	// Process the received JSON data as needed
	// ...

	response := jsonData.(map[string]interface{})

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("Error creating JSON response:", err)
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
func getEmbedDetails(w http.ResponseWriter, r *http.Request) {
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
			serverAPIUrl := queryString.(map[string]interface{})["dashboardServerApiUrl"].(string)
			embedQueryString := queryString.(map[string]interface{})["embedQuerString"].(string)
			embedQueryString += "&embed_user_email=" + userMail
			timeStamp := time.Now().Unix()
			embedQueryString += "&embed_server_timestamp=" + strconv.FormatInt(timeStamp, 10)
			signatureString, err := getSignatureUrl(embedQueryString)
			embedDetails := "/embed/authorize?" + embedQueryString + "&embed_signature=" + signatureString
			query := serverAPIUrl + embedDetails
			log.Println(query)
			result, err := http.Get(query)
			if err != nil {
				log.Println(err)
				http.Error(w, "Error calling API", http.StatusInternalServerError)
				return
			}
			log.Println(result)
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
