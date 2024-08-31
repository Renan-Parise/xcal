package analytes

import "time"

type Values struct {
	Values []Information `json:"values"`
}

type Information struct {
	Value      []float64 `json:"value"`
	IncludedAt time.Time `json:"included_at"`
}

type Analytes struct {
	Glucose                   Values `json:"glucose"`
	Hemoglobin                Values `json:"hemoglobin"`
	Cholesterol               Values `json:"cholesterol"`
	Creatinine                Values `json:"creatinine"`
	BloodUreaNitrogen         Values `json:"blood_urea_nitrogen"`
	WhiteBloodCellCount       Values `json:"white_blood_cell_count"`
	Sodium                    Values `json:"sodium"`
	Potassium                 Values `json:"potassium"`
	AlanineAminotransferase   Values `json:"alanine_aminotransferase"`
	AspartateAminotransferase Values `json:"aspartate_aminotransferase"`
	Calcium                   Values `json:"calcium"`
	Magnesium                 Values `json:"magnesium"`
	ThyroidStimulatingHormone Values `json:"thyroid_stimulating_hormone"`
	CReactiveProtein          Values `json:"c_reactive_protein"`
}
