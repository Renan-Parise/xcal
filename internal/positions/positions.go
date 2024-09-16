package positions

import "time"

type Position struct {
	ID           int           `json:"id"`
	Title        string        `json:"title"`
	Company      string        `json:"company"`
	Location     string        `json:"location"`
	Description  string        `json:"description"`
	Salary       float64       `json:"salary"`
	Currency     string        `json:"currency"`
	IsRemote     bool          `json:"isRemote"`
	Requirements []Requirement `json:"requirements"`
	Benefits     []Benefit     `json:"benefits"`
	Start        time.Time     `json:"startDate"`
	End          time.Time     `json:"endDate"`
	Created      time.Time     `json:"postedDate"`
}

type Requirement struct {
	Name       string `json:"name"`
	Obligatory bool   `json:"obligatory"`
}

type Benefit struct {
	Name string `json:"name"`
}
