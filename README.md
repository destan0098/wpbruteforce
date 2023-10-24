# Find subdomains And Check Port 80 and 443 Open Or not And check CDNs
![Static Badge](https://img.shields.io/badge/Go-100%25-brightgreen)
## Description

This Tool Check Subdomains has Open Port 80 or 443 and use cdn or not
This tool is for training.


- Check http and https
- Show CDN name
- Find Subdomains
 Help You to find real ip


## Table of Contents 


- [Installation](#installation)
- [Usage](#usage)


## Installation

```
go install github.com/destan0098/subenum/cmd/subenum@latest
```
or use
```
git clone https://github.com/destan0098/subenum.git

```

## Usage

To Run Enter Below Code
For Use This Enter Website without http  In Input File
Like : google.com

```
subcheck -l 'C:\Users\**\Desktop\go2\checksubdomains\input.txt' -o 'C:\Users\***\Desktop\go2\checksubdomains\result4.csv'

```
```
subcheck -d google.com 
```
```
cat inputfile.txt | subcheck -pipe -o output.csv
```
```
NAME:
   subcheck - A new cli application

USAGE:
   subcheck [global options] command [command options] [arguments...]

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

## ScreenShot

![IP Show](/screenshot2.png?raw=true "subenum")


## Features

This Tool Check and find Subdomains has Open Port 80 or 443 and use cdn or not


