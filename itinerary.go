package main

import (
    "encoding/csv"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "regexp"
    "strings"
    "time"
)

func main() {
    // Define and parse command line arguments
    help := flag.Bool("h", false, "Display usage")
    flag.Parse()

    // Display usage if -h flag is provided or incorrect number of arguments are provided
    if *help || len(flag.Args()) != 3 {
        fmt.Println("itinerary usage:\ngo run . ./input.txt ./output.txt ./airport-lookup.csv")
        return
    }

    // Get the paths from the arguments
    inputPath := flag.Arg(0)
    outputPath := flag.Arg(1)
    lookupPath := flag.Arg(2)

    // Check if the input file exists
    if _, err := os.Stat(inputPath); os.IsNotExist(err) {
        fmt.Println("Input not found")
        return
    }

    // Check if the airport lookup file exists
    if _, err := os.Stat(lookupPath); os.IsNotExist(err) {
        fmt.Println("Airport lookup not found")
        return
    }

    // Open the airport lookup file
    csvfile, err := os.Open(lookupPath)
    if err != nil {
        log.Fatalln("Couldn't open the csv file", err)
    }

    // Parse the airport lookup file
    r := csv.NewReader(csvfile)
    lookupData, err := r.ReadAll()
    if err != nil {
        log.Fatalln("Failed to parse the csv file", err)
    }

    // Create a map for easy lookup of airport codes and city names
    lookupMap := make(map[string]string)
    cityLookupMap := make(map[string]string)
    for _, record := range lookupData {
        if record[0] == "" || record[2] == "" || record[3] == "" || record[4] == "" {
            fmt.Println("Airport lookup malformed")
            return
        }
        lookupMap[record[3]] = record[0] // ICAO code to airport name
        lookupMap[record[4]] = record[0] // IATA code to airport name
        cityLookupMap[record[3]] = record[2] // ICAO code to city name
        cityLookupMap[record[4]] = record[2] // IATA code to city name
    }

    // Read the input file
    inputData, err := ioutil.ReadFile(inputPath)
    if err != nil {
        log.Fatalln("Failed to read the input file", err)
    }

    // Convert the input data to a string
    inputString := string(inputData)

    // Replace line-break characters with new-line characters
    inputString = strings.ReplaceAll(inputString, "\v", "\n")
    inputString = strings.ReplaceAll(inputString, "\f", "\n")
    inputString = strings.ReplaceAll(inputString, "\r", "\n")

    // Replace multiple consecutive new-line characters with a single one
    re := regexp.MustCompile("\n{3,}")
    inputString = re.ReplaceAllString(inputString, "\n\n")

    // Process the input data
    outputData := inputString
    for code, city := range cityLookupMap {
        if _, exists := cityLookupMap[code]; exists {
            outputData = strings.ReplaceAll(outputData, "*##"+code, city)
            outputData = strings.ReplaceAll(outputData, "*#"+code, city)
        }
    }
    for code, name := range lookupMap {
        if _, exists := lookupMap[code]; exists {
            outputData = strings.ReplaceAll(outputData, "##"+code, name)
            outputData = strings.ReplaceAll(outputData, "#"+code, name)
        }
    }

    // Define the date and time formats
    const isoFormat = "2006-01-02T15:04-07:00"
    const dateFormat = "02-Jan-2006"
    const time12Format = "03:04PM (-07:00)"
    const time24Format = "15:04 (-07:00)"

    // Define the regular expressions for the date and time patterns
    reDate := regexp.MustCompile(`D\((.+?)\)`)
    reTime12 := regexp.MustCompile(`T12\((.+?)\)`)
    reTime24 := regexp.MustCompile(`T24\((.+?)\)`)

    // Replace the date and time patterns with the customer friendly format
    outputData = reDate.ReplaceAllStringFunc(outputData, func(s string) string {
        s = strings.Replace(s, "−", "-", -1) // Replace the minus sign with a hyphen-minus
        t, err := time.Parse(isoFormat, s[2:len(s)-1])
        if err != nil {
            return s // If the date is malformed, leave it unchanged
        }
        return t.Format(dateFormat)
    })
    outputData = reTime12.ReplaceAllStringFunc(outputData, func(s string) string {
        s = strings.Replace(s, "−", "-", -1) // Replace the minus sign with a hyphen-minus
        t, err := time.Parse(isoFormat, s[4:len(s)-1])
        if err != nil {
            return s // If the time is malformed, leave it unchanged
        }
        return t.Format(time12Format)
    })
    outputData = reTime24.ReplaceAllStringFunc(outputData, func(s string) string {
        s = strings.Replace(s, "−", "-", -1) // Replace the minus sign with a hyphen-minus
        t, err := time.Parse(isoFormat, s[4:len(s)-1])
        if err != nil {
            return s // If the time is malformed, leave it unchanged
        }
        return t.Format(time24Format)
    })

    // Write the processed data to the output file
    err = ioutil.WriteFile(outputPath, []byte(outputData), 0644)
    if err != nil {
        log.Fatalln("Failed to write the output file", err)
    }
}