package comscape

type DAOFake struct {
}

func (d DAOFake) GetMatrixData(domain string) (map[string]string, error) {
	return map[string]string{}, nil
}

// getDNData gathers data from Datanyze and unmarshals it into a struct
func (d DAOFake) GetDNData(domain string) (dnResponse, error) {
	return dnResponse{
		Domain:              "tune.com",
		Rank:                65921,
		Phone:               "+12065081318",
		Email:               "hello@tune.com",
		CompanyName:         "Tune, Inc",
		Country:             236,
		State:               48,
		City:                "Seattle",
		Street:              "2200 Western Avenue",
		FullAddress:         "2200 Western Avenue Seattle WA 98121 United States",
		Employees:           4,
		Revenue:             4,
		Funded:              1,
		Public:              2,
		IndustryID:          84,
		Twitter:             "tune",
		SicCode:             7311,
		NaicsCode:           541613,
		FoundedYear:         2009,
		TotalMoneyRaised:    "36.4M",
		CountryName:         "United States",
		CountryIso:          "USA",
		StateName:           "WA",
		TotalMoneyRaisedInt: 36400000,
		EmployeesStr:        "250 - 1000",
		RevenueStr:          "$50 - 100M",
		IndustryName:        "Marketing and Advertising",
		TwitterFollowers:    15276,
		SocialLinks: socialStruct{
			Linkedin:   "http://www.linkedin.com/company/635323",
			Facebook:   "https://www.facebook.com/tune",
			Google:     "https://plus.google.com/110653481875028879836",
			Blog:       "http://www.tune.com/blog",
			Crunchbase: "https://crunchbase.com/organization/tune",
			Angellist:  "https://angel.co/tune",
			Twitter:    "https://twitter.com/tune"},
		MonthlyTechSpend: "$20K - 50K", Tags: []string{"sf_accounts", "SFDC", "sf_leads"}}, nil
}
