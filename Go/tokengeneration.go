package main

import (
    "encoding/json"
    "os"
    "io"
    "bytes"
    "log"
    "net/http"
    "strings"
    "fmt"
)

var embedConfig map[string]interface{}

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
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    data, err := os.ReadFile("embedConfig.json")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "embedConfig.json file not found"})
        return
    }

    // trim BOM only and unmarshal (don't strip '//' since URLs contain it)
    s := strings.TrimPrefix(string(data), "\uFEFF")
    var cfg map[string]interface{}
    if err := json.Unmarshal([]byte(strings.TrimSpace(s)), &cfg); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "invalid embedConfig.json: " + err.Error()})
        return
    }
    // persist for other handlers
    embedConfig = cfg

    // helper to safely get string fields
    getStr := func(m map[string]interface{}, key string) string {
        if v, ok := m[key]; ok && v != nil {
            if s, ok := v.(string); ok {
                return s
            }
        }
        return ""
    }

    clientEmbedConfigData := EmbedConfig{
        DashboardId:    getStr(cfg, "DashboardId"),
        ServerUrl:      getStr(cfg, "ServerUrl"),
        SiteIdentifier: getStr(cfg, "SiteIdentifier"),
        EmbedType:      getStr(cfg, "EmbedType"),
        Environment:    getStr(cfg, "Environment"),
    }

    // Validate required fields
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
    // CORS headers
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    // ensure embedConfig is loaded
    if embedConfig == nil {
        data, err := os.ReadFile("embedConfig.json")
        if err != nil {
            http.Error(w, "embedConfig.json file not found", http.StatusInternalServerError)
            return
        }
        var cfg map[string]interface{}
        if err := json.Unmarshal(data, &cfg); err != nil {
            http.Error(w, "invalid embedConfig.json", http.StatusInternalServerError)
            return
        }
        embedConfig = cfg
    }

    // helper to safely read string fields without repeating .(string)
    getStr := func(m map[string]interface{}, keys ...string) string {
        for _, k := range keys {
            if v, ok := m[k]; ok && v != nil {
                if s, ok := v.(string); ok {
                    return s
                }
            }
        }
        return ""
    }

    serverUrl := getStr(embedConfig, "ServerUrl", "serverurl")
    siteIdentifier := getStr(embedConfig, "SiteIdentifier", "siteidentifier")
    email := getStr(embedConfig, "UserEmail", "email")
    embedSecret := getStr(embedConfig, "EmbedSecret", "embedsecret")
    dashboardId := getStr(embedConfig, "DashboardId", "dashboardId")

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

    payloadBytes, err := json.Marshal(embedDetails)
    if err != nil {
        http.Error(w, "failed to marshal embedDetails", http.StatusInternalServerError)
        return
    }

    requestUrl := fmt.Sprintf("%s/api/%s/embed/authorize", serverUrl, siteIdentifier)
    resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(payloadBytes))
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

    // Try to parse JSON response; if not JSON, return raw body
    var respObj map[string]interface{}
    if err := json.Unmarshal(respBytes, &respObj); err != nil {
        log.Println("authorize response (non-JSON):", string(respBytes))
        w.Write(respBytes)
        return
    }

    // helper to extract access_token from common places
    findToken := func(obj map[string]interface{}) string {
        // check Data or data first
        for _, key := range []string{"Data", "data"} {
            if v, ok := obj[key]; ok {
                if m, ok := v.(map[string]interface{}); ok {
                    if t, ok := m["access_token"].(string); ok {
                        return t
                    }
                }
            }
        }
        // check top-level
        if t, ok := obj["access_token"].(string); ok {
            return t
        }
        return ""
    }

    token := findToken(respObj)
    if token == "" {
        http.Error(w, "access_token not found in response", http.StatusBadGateway)
        return
    }

    w.Write([]byte(token))
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