# Find subdomains And Check Port 80 and 443 Open Or not And check CDNs
![Static Badge](https://img.shields.io/badge/Go-100%25-brightgreen)
## Description

This Tool BruteForce Wordpress with xmlrpc
This tool is for training.





## Table of Contents 


- [Installation](#installation)
- [Usage](#usage)


## Installation

```
go install github.com/destan0098/wpbruteforce/cmd/wpbruteforce@latest
```
or use
```
git clone https://github.com/destan0098/wpbruteforce.git

```

## Usage

To Run Enter Below Code
For Use This Enter Website without http  In Input File
Like : google.com

```
wpbruteforce -l 'input.txt' -o 'result4.csv' -u username.txt -w password.txt

```
```
wpbruteforce -d google.com  -u username.txt -w password.txt
```
```
cat inputfile.txt | wpbruteforce -pipe -o output.csv -u username.txt -w password.txt
```
```


USAGE:
   wpbruteforce [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --domain value, -d value  Enter just one domain
   --list value, -l value    Enter a list from text file
   --pipe                    Enter just from pipe line (default: false)
   --output value, -o value  Enter output csv file name   (default: "output.csv")
   --help, -h                show help

```




---


## Features

This Tool BruteForce Wordpress with xmlrpc


