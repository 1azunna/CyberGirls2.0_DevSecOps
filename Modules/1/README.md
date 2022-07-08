# Module 1 -  Introduction

In this week's module, we created a snyk parser in 3 different languages which takes a snyk report as an input argument.

## Set up

Ensure the following are installed.
- Go v1.18 or higher
- jq
- python3.8 or higher

## Bash Development
To run the script:
```
./snyk-parser.sh snyk-report.json
```

## Python Development
Ensure the dependent libraries are installed by running `pip install -r requirements.txt`

To run the package:
```
python -m snyk-parser snyk-report.json
```

## Go Development
To run the package:
```
go run ./snyk-parser.go snyk-report.json
```