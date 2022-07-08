# Import the necessary libraries
import json
import sys
from prettytable import PrettyTable

# Assign the filename passed as an argument to a variable
filename = sys.argv[1]

# Open the file, read the contents and save it to a variable 'file' then load the data as a JSON dictionary.
# Using 'with' statement implements the file.close() method which is a good practice for saving system resources.
with open(filename, 'r') as file:
    data = json.load(file)

# Create the Table and Table headers.
table = PrettyTable()
table.field_names = ["Title", "Severity", "CVSS Score"]

# You want to ensure that there is data on the file before trying to manipulate it.
if len(data) > 0:
    # Get all vulnerabilities using the appropriate dict key.
    vulnerabilities = data['vulnerabilities']

    # Loop through vulnerabilities to find the title, severity and cvss score values. Append those values to the table.
    for finding in vulnerabilities:
        title = finding['title']
        severity = finding['severity']
        cvss_score = finding['cvssScore']
        table.add_row([title, severity, cvss_score])

table.title = "Snyk Scan Results" # Set the title of the table
table.sortby = "CVSS Score" # Sort the table by CVSS score
table.reversesort = True # Sort the table in descending order.
print(table.get_string()) # Print the table