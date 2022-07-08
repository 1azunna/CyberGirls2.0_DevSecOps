#!/bin/sh

# Check if filename is specified
if [ -z "$1" ];
then
    echo "Filename not specified";
    exit 1;
fi

# Assign the filename to a variable
filename=$1

# Use jq utility to create a tab-separated output data with headers - "Title", "Severity" and"CVSS Score"
# Use the expand command to expand the output data, using a tab spacing of 65.
jq -r '["Title","Severity","CVSS Score"], (.vulnerabilities[] | [.title, .severity, .cvssScore]) | @tsv' $filename | expand -t 65