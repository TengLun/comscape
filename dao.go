package comscape

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type DAO struct {
}

// GetMatrixData gathers data from the competitive matrix for a given partner
func (d DAO) GetMatrixData(domain string) (map[string]string, error) {
	csvFile, _ := ioutil.ReadFile("matrix.csv")
	csvFileReader := strings.NewReader(string(csvFile))
	f, _ := csv.NewReader(csvFileReader).ReadAll()
	for i := range f {
		for ii := range f[i] {
			fmt.Print(f[i][ii] + ", ")
		}
		fmt.Print("\n")
	}
	return map[string]string{}, nil
}

// getDNData gathers data from Datanyze and unmarshals it into a struct
func (d DAO) GetDNData(domain string) (dnResponse, error) {
	url := `https://api.datanyze.com/domain_info/?email=jason%40kochava.com&token=e1f3ca2fb3888858dc697657e7439e63&domain=` + domain

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return dnResponse{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return dnResponse{}, err
	}

	var dnr dnResponse

	err = json.Unmarshal(body, &dnr)
	if err != nil {
		fmt.Println(err)
		return dnResponse{}, err
	}

	return dnr, nil
}
