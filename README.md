# Itinerary Processor

This is a Go program that processes an itinerary file, replacing airport codes with their corresponding names and city names, and converting date and time formats to a more readable format.

## Usage

To run the program, use the following command:

 go run . ./input.txt ./output.txt ./airport-lookup.csv

The program takes three arguments:

1. input.txt: The path to the input file. This file should contain the itinerary data with airport codes and dates/times in a specific format.
2. output.txt: The path where the output file will be written. This file will contain the processed itinerary data with airport codes replaced by their corresponding names and city names, and dates/times converted to a more readable format.
3. airport-lookup.csv: The path to a CSV file that maps airport codes to their corresponding names and city names.

## Input Format

The input file should contain itinerary data with airport codes and dates/times in the following format:

- Airport codes are represented as ##CODE or #CODE for ICAO and IATA codes respectively.
- City names are represented as *##CODE or *#CODE for ICAO and IATA codes respectively.
- Dates are represented as D(YYYY-MM-DDTHH:MM-SS:00).
- Times are represented as T12(HH:MM-SS:00) for 12-hour format and T24(HH:MM-SS:00) for 24-hour format.

## Airport Lookup Format

The airport lookup file should be a CSV file with each record in the following format:

- Column 1: Airport name
- Column 2: City name
- Column 3: ICAO code
- Column 4: IATA code

## Output Format

The output file will contain the processed itinerary data with airport codes replaced by their corresponding names and city names, and dates/times converted to a more readable format:

Airport codes are replaced with their corresponding names.
City names are replaced with their corresponding names.
Dates are converted to DD-MMM-YYYY format.
Times are converted to HH:MMAM/PM (-SS:00) for 12-hour format and HH:MM (-SS:00) for 24-hour format.

## Error Handling

The program will terminate with an error message if:

- The input file or airport lookup file does not exist.
- The airport lookup file is malformed.
- There is an error reading the input file or writing the output file.