package analytes

func (a *Analytes) Merge(newAnalytes Analytes) {
	merge := func(existing, new Values) Values {
		if existing.Values == nil {
			return new
		}
		existing.Values = append(existing.Values, new.Values...)
		return existing
	}

	a.Glucose = merge(a.Glucose, newAnalytes.Glucose)
	a.Hemoglobin = merge(a.Hemoglobin, newAnalytes.Hemoglobin)
	a.Cholesterol = merge(a.Cholesterol, newAnalytes.Cholesterol)
	a.Creatinine = merge(a.Creatinine, newAnalytes.Creatinine)
	a.BloodUreaNitrogen = merge(a.BloodUreaNitrogen, newAnalytes.BloodUreaNitrogen)
	a.WhiteBloodCellCount = merge(a.WhiteBloodCellCount, newAnalytes.WhiteBloodCellCount)
	a.Sodium = merge(a.Sodium, newAnalytes.Sodium)
	a.Potassium = merge(a.Potassium, newAnalytes.Potassium)
	a.AlanineAminotransferase = merge(a.AlanineAminotransferase, newAnalytes.AlanineAminotransferase)
	a.AspartateAminotransferase = merge(a.AspartateAminotransferase, newAnalytes.AspartateAminotransferase)
	a.Calcium = merge(a.Calcium, newAnalytes.Calcium)
	a.Magnesium = merge(a.Magnesium, newAnalytes.Magnesium)
	a.ThyroidStimulatingHormone = merge(a.ThyroidStimulatingHormone, newAnalytes.ThyroidStimulatingHormone)
	a.CReactiveProtein = merge(a.CReactiveProtein, newAnalytes.CReactiveProtein)
}
