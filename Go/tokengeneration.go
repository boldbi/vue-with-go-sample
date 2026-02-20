package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "bytes"
    "strings"
    "sync"
)

var embedConfig map[string]interface{}
var embedConfigOnce sync.Once
var embedConfigErr error

type EmbedConfig struct {
    DashboardId    string `json:"DashboardId"`
    ServerUrl      string `json:"ServerUrl"`
    EmbedType      string `json:"EmbedType"`
    Environment    string `json:"Environment"`
    SiteIdentifier string `json:"SiteIdentifier"`
}

func main() {
    http.HandleFunc("/tokenGeneration", tokenGeneration)
    http.HandleFunc("/getdetails", getdetails)
    fmt.Println("Go server is running on port 8086")
    log.Fatal(http.ListenAndServe(":8086", nil))
}

func getdetails(w http.ResponseWriter, r *http.Request) {
    setCORS(w, "GET")

    if err := loadEmbedConfig(); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }

    clientEmbedConfigData := EmbedConfig{
        DashboardId:    getConfigStr("DashboardId"),
        ServerUrl:      getConfigStr("ServerUrl"),
        SiteIdentifier: getConfigStr("SiteIdentifier"),
        EmbedType:      getConfigStr("EmbedType"),
        Environment:    getConfigStr("Environment"),
    }

    if clientEmbedConfigData.ServerUrl == "" || clientEmbedConfigData.SiteIdentifier == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "ServerUrl and SiteIdentifier are required in embedConfig.json"})
        return
    }

    jsonResponse, err := json.Marshal(clientEmbedConfigData)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "failed to marshal response"})
        return
    }
    w.Write(jsonResponse)
}

func tokenGeneration(w http.ResponseWriter, r *http.Request) {
    setCORS(w, "POST, OPTIONS")

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    if err := loadEmbedConfig(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    serverUrl := getConfigStr("ServerUrl", "serverurl")
    siteIdentifier := getConfigStr("SiteIdentifier", "siteidentifier")
    email := getConfigStr("UserEmail", "email")
    embedSecret := getConfigStr("EmbedSecret", "embedsecret")
    dashboardId := getConfigStr("DashboardId", "dashboardId")

    if serverUrl == "" || siteIdentifier == "" {
        http.Error(w, "ServerUrl and SiteIdentifier are required", http.StatusBadRequest)
        return
    }

    embedDetails := map[string]interface{}{
        "serverurl":      serverUrl,
        "siteidentifier": siteIdentifier,
        "email":          email,
        "embedsecret":    embedSecret,
        "dashboard": map[string]string{
            "id": dashboardId,
        },
    }

    payload, err := json.Marshal(embedDetails)
    if err != nil {
        http.Error(w, "failed to marshal embedDetails", http.StatusInternalServerError)
        return
    }

    requestUrl := fmt.Sprintf("%s/api/%s/embed/authorize", strings.TrimRight(embedDetails["serverurl"].(string), "/"), embedDetails["siteidentifier"].(string))
    resp, err := http.Post(requestUrl, "application/json", bytes.NewReader(payload))
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadGateway)
        return
    }
    defer resp.Body.Close()

    respBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, "failed to read response", http.StatusBadGateway)
        return
    }

    var respObj map[string]interface{}
    if err := json.Unmarshal(respBytes, &respObj); err != nil {
        w.WriteHeader(resp.StatusCode)
        w.Write(respBytes)
        return
    }

    d, ok := respObj["Data"].(map[string]interface{})
    if !ok {
        http.Error(w, "invalid response shape: Data field missing", http.StatusBadGateway)
        return
    }
    token, ok := d["access_token"].(string)
    if !ok || token == "" {
        http.Error(w, "access_token not found in Data", http.StatusBadGateway)
        return
    }
    w.Write([]byte(token))
}

// loadEmbedConfig reads and parses embedConfig.json once and caches the result.
func loadEmbedConfig() error {
    embedConfigOnce.Do(func() {
        data, err := os.ReadFile("embedConfig.json")
        if err != nil {
            embedConfigErr = fmt.Errorf("embedConfig.json file not found: %w", err)
            return
        }
        s := strings.TrimPrefix(string(data), "\uFEFF")
        var cfg map[string]interface{}
        if err := json.Unmarshal([]byte(strings.TrimSpace(s)), &cfg); err != nil {
            embedConfigErr = fmt.Errorf("invalid embedConfig.json: %w", err)
            return
        }
        embedConfig = cfg
    })
    return embedConfigErr
}

// setCORS writes common CORS and JSON headers
func setCORS(w http.ResponseWriter, methods string) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", methods)
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// getConfigStr returns the first non-empty string value from embedConfig for provided keys
func getConfigStr(keys ...string) string {
    for _, k := range keys {
        if s, ok := embedConfig[k].(string); ok && s != "" {
            return s
        }
    }
    return ""
}