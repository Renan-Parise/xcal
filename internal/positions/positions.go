package positions

import "time"

type Metric any

type Values struct {
	Values []Information `json:"values"`
}

type Information struct {
	Value      Metric    `json:"value"`
	IncludedAt time.Time `json:"included_at"`
}

type Positions struct {
	AlanineAminotransferase   Values `bson:"alanine_aminotransferase" json:"alanine_aminotransferase"`
	AspartateAminotransferase Values `bson:"aspartate_aminotransferase" json:"aspartate_aminotransferase"`
	BloodUreaNitrogen         Values `bson:"blood_urea_nitrogen" json:"blood_urea_nitrogen"`
	Calcium                   Values `bson:"calcium" json:"calcium"`
	Cholesterol               Values `bson:"cholesterol" json:"cholesterol"`
	Creatinine                Values `bson:"creatinine" json:"creatinine"`
	CReactiveProtein          Values `bson:"c_reactive_protein" json:"c_reactive_protein"`
	Glucose                   Values `bson:"glucose" json:"glucose"`
	Hemoglobin                Values `bson:"hemoglobin" json:"hemoglobin"`
	Magnesium                 Values `bson:"magnesium" json:"magnesium"`
	Potassium                 Values `bson:"potassium" json:"potassium"`
	Sodium                    Values `bson:"sodium" json:"sodium"`
	ThyroidStimulatingHormone Values `bson:"thyroid_stimulating_hormone" json:"thyroid_stimulating_hormone"`
	WhiteBloodCellCount       Values `bson:"white_blood_cell_count" json:"white_blood_cell_count"`
}
