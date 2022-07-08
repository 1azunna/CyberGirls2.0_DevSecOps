package main //Define the main package. This is required by go to run the package.

// Import the necessary libraries
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/olekukonko/tablewriter"
)

// We define our data structure here which matches the input data structure and types. This is used by the json library to parse our json data.
type Findings struct {
	Title     string  `json:"title"`
	Severity  string  `json:"severity"`
	CvssScore float64 `json:"cvssScore"`
}

type Vulnerabilities struct {
	Vulnerabilities []Findings
}

// The main function which is what is executed.
func main() {

	// Assign the filename passed as an argument to a variable
	filename := os.Args[1]

	// Open the file, read the contents and save it to a variable 'data' of type []byte.
	data, err := ioutil.ReadFile(filename)
	// handle any errors that occur when opening the file.
	if err != nil {
		fmt.Printf("Error reading file")
		os.Exit(1)
	}
	// Define the variables that will hold the unmarshalled json data, and output data for the table which is a list of (list of strings) [][]string.
	var jsonData Vulnerabilities
	var output [][]string
	// Create the Table and Table headers.
	table := tablewriter.NewWriter(os.Stdout)                    // Create a table writer that will send output to Stdout
	table.SetHeader([]string{"Title", "Severity", "CVSS Score"}) // Define the headers for the table

	// unmarshall the json data into the previously defined variable of appropriate data type.
	json.Unmarshal(data, &jsonData)
	// fmt.Printf("%v\n", jsonData) //Uncomment if you want to print out the unmarshalled JSON data

	// Loop through the vulnerabilities to find the title, severity and cvss score values. Append those values to the table.
	for i := 0; i < len(jsonData.Vulnerabilities); i++ {
		title := jsonData.Vulnerabilities[i].Title
		severity := jsonData.Vulnerabilities[i].Severity
		cvss_score := jsonData.Vulnerabilities[i].CvssScore
		output = append(output, []string{title, severity, fmt.Sprint(cvss_score)})
	}

	table.AppendBulk(output)                   // AppendBulk function helps append the output data which is a list of list of strings. Alternatively, you can append data to the table when looping.
	table.SetAlignment(tablewriter.ALIGN_LEFT) // Set alignment for the table data to the left.
	table.Render()                             // Print the output

}
