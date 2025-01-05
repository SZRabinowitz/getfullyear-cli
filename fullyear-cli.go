package main

import (
        "encoding/json"
        "flag"
        "fmt"
        "net/http"
)

const apiURL = "https://getfullyear.com/api/year"

type YearResponse struct {
        Year        int    `json:"year"`
        YearString  string `json:"year_string"`
        SponsoredBy string `json:"sponsored_by"`
}

func main() {
        // Define flags
        jsonOutput := flag.Bool("json", false, "Output in JSON format")
        copyright := flag.String("copyright", "", "Print copyright message with the specified value")
        flag.Parse()

        // Fetch and display year data
        fetchYearData(*jsonOutput, *copyright)
}

func fetchYearData(jsonOutput bool, copyright string) {
        resp, err := http.Get(apiURL)
        if err != nil {
                fmt.Printf("Error fetching data: %s\n", err)
                return
        }
        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
                fmt.Printf("Error: received status code %d\n", resp.StatusCode)
                return
        }

        var data YearResponse
        if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
                fmt.Printf("Error decoding response: %s\n", err)
                return
        }

        console := struct {
                log func(string)
        }{
                log: func(message string) {
                        fmt.Println(message)
                },
        }

        console.log(data.SponsoredBy)

        if jsonOutput {
                output, err := json.MarshalIndent(data, "", "  ")
                if err != nil {
                        fmt.Printf("Error formatting JSON: %s\n", err)
                        return
                }
                fmt.Println(string(output))
        } else {
                fmt.Println("\n--- Year Data ---")
                fmt.Printf("Year: %d\n", data.Year)
                fmt.Printf("Year String: %s\n", data.YearString)
                //fmt.Printf("Sponsored By: %s\n", data.SponsoredBy)
        }

        if copyright != "" {
                fmt.Printf("\nCopyright © %d %s\n", data.Year, copyright)
        }
}
