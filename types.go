package comscape

type DataAccessor interface {
	GetDNData(domain string) (dnResponse, error)
	GetMatrixData(domain string) (map[string]string, error)
}

type dnResponse struct {
	Domain              string `json:"domain"`
	Rank                int    `json:"rank"`
	Phone               string `json:"phone"`
	Email               string `json:"email"`
	CompanyName         string `json:"company_name"`
	Country             int    `json:"country"`
	State               int    `json:"state"`
	City                string `json:"city"`
	Street              string `json:"street"`
	FullAddress         string `json:"full_address"`
	Employees           int    `json:"employees"`
	Revenue             int    `json:"revenue"`
	Funded              int    `json:"funded"`
	Public              int    `json:"public"`
	IndustryID          int    `json:"industry_id"`
	Twitter             string `json:"twitter"`
	SicCode             int    `json:"sic_code"`
	NaicsCode           int    `json:"naics_code"`
	FoundedYear         int    `json:"founded_year"`
	TotalMoneyRaised    string `json:"total_money_raised"`
	CountryName         string `json:"country_name"`
	CountryIso          string `json:"country_iso"`
	StateName           string `json:"state_name"`
	TotalMoneyRaisedInt int    `json:"total_money_raised_int"`
	EmployeesStr        string `json:"employees_str"`
	RevenueStr          string `json:"revenue_str"`
	IndustryName        string `json:"industry_name"`
	TwitterFollowers    int    `json:"twitter_followers"`
	SocialLinks         struct {
		Linkedin   string `json:"linkedin"`
		Facebook   string `json:"facebook"`
		Google     string `json:"google"`
		Blog       string `json:"blog"`
		Crunchbase string `json:"crunchbase"`
		Angellist  string `json:"angellist"`
		Twitter    string `json:"twitter"`
	} `json:"social_links"`
	MonthlyTechSpend string   `json:"monthly_tech_spend"`
	Tags             []string `json:"tags"`
}

type socialStruct struct {
	Linkedin   string `json:"linkedin"`
	Facebook   string `json:"facebook"`
	Google     string `json:"google"`
	Blog       string `json:"blog"`
	Crunchbase string `json:"crunchbase"`
	Angellist  string `json:"angellist"`
	Twitter    string `json:"twitter"`
}
