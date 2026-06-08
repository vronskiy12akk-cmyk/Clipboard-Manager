package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "time"
    "github.com/atotto/clipboard"
)

type Entry struct {
    Timestamp string `json:"timestamp"`
    Content   string `json:"content"`
}

const historyFile = "history.json"

func loadHistory() []Entry {
    data, err := ioutil.ReadFile(historyFile)
    if err != nil {
        return []Entry{}
    }
    var history []Entry
    json.Unmarshal(data, &history)
    return history
}

func saveHistory(history []Entry) {
    data, _ := json.MarshalIndent(history, "", "  ")
    ioutil.WriteFile(historyFile, data, 0644)
}

func main() {
    history := loadHistory()
    last, _ := clipboard.ReadAll()
    fmt.Println("Clipboard Manager запущен. Ctrl+C для выхода.")
    for {
        current, _ := clipboard.ReadAll()
        if current != last && current != "" {
            entry := Entry{
                Timestamp: time.Now().Format(time.RFC3339),
                Content:   current,
            }
            history = append(history, entry)
            saveHistory(history)
            fmt.Printf("[%s] Скопировано: %s...\n", time.Now().Format("15:04:05"), truncate(current, 50))
            last = current
        }
        time.Sleep(1 * time.Second)
    }
}

func truncate(s string, n int) string {
    if len(s) <= n {
        return s
    }
    return s[:n]
}
