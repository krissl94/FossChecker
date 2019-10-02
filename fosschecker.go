package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Person struct {
	StudentNumber       string `json:"StudentNumber"`
	FullName            string `json:"FullName"`
	Email               string `json:"Email"`
	UserName            string `json:"UserName"`
	ProfileImageAddress string `json:"ProfileImageAddress"`
	Department          string `json:"Department"`
	AadObjectId         string `json:"AadObjectId"`
	UserPrincipleName   string `json:"UserPrincipleName"`
	ItemSource          string `json:"ItemSource"`
}

func (p Person) toCsv() string {
	return p.StudentNumber + "," +
		p.FullName + "," +
		p.Email + "," +
		p.UserName + "," +
		p.ProfileImageAddress + ",\"" +
		p.Department + "\"," +
		p.AadObjectId + "," +
		p.UserPrincipleName + "," +
		p.ItemSource
}
func getJSON(studentNr string, cookie string) (Person, error) {
	pers := Person{}

	url := "https://eur.delve.office.com/mt/v3/people/" + studentNr + "%40student.saxion.nl"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cookie", "X-Delve-AuthEur="+cookie)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "eur.delve.office.com")
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("Accept-Encoding", "")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return pers, err
	}
	defer res.Body.Close()				//Close the body when the function is done. 
	body, _ := ioutil.ReadAll(res.Body)

	pers.StudentNumber = studentNr
	json.Unmarshal(body, &pers)
	return pers, nil
}
func selectFile() *os.File {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Type your input csv file: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%q", text)

	text = strings.TrimSuffix(text, "\n")
	fmt.Println(text)

	csvFile, err := os.Open(text)
	if err != nil {
		//Can't open that file. Ask again.
		fmt.Println("Errored on that.. Is your path correct?")
		log.Fatal(err)
		selectFile()
	}
	return csvFile
}
func main() {
	//Ask for input file
	file := selectFile()
	fileReader := csv.NewReader(bufio.NewReader(file))
	//Ask for cookie
	//The cookie can be retrieved by logging in to eur.delve.office.com and copying the value in the X-Delve-AuthEur cookie
	//Use the Application tab in Chrome's devTools to read cookies. 
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("We'll need your X-Delve-AuthEur cookie value from eur.delve.office.com. Please copy the value here: ")
	cookie, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error on the cookie")
		log.Fatal(err)
	}

	cookie = strings.TrimSuffix(cookie, "\n")
	
	outputFile, err := os.Create("output.csv")

	_, err = outputFile.WriteString("StudentNumber,FullName,Email,UserName,ProfileImageAddress,Department,AadObjectId,UserPrincipleName,ItemSource\r\n")

	if err != nil {
		fmt.Println("Can't write to output.csv, is it locked?")
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal("Couldn't create output.csv. Does it already exist?")
		log.Fatal(err)
	}
	defer outputFile.Close()
	defer file.Close()

	for {
		line, error := fileReader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		fmt.Println(line)
		// Run each student number through the who-is-who
		pers, jsonErr := getJSON(line[0], cookie)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		values := pers.toCsv()

		_, err := outputFile.WriteString(values + "\r\n")

		if err != nil {
			//write failed do something
			log.Fatal(err)
		}

	}
	fmt.Println("Output: output.csv")
}
