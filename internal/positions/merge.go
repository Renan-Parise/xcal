package positions

func (a *Positions) Merge(newPositions Positions) {
	merge := func(existing, new Values) Values {
		if existing.Values == nil {
			return new
		}
		existing.Values = append(existing.Values, new.Values...)
		return existing
	}

	a.Glucose = merge(a.Glucose, newPositions.Glucose)
	a.Hemoglobin = merge(a.Hemoglobin, newPositions.Hemoglobin)
	a.Cholesterol = merge(a.Cholesterol, newPositions.Cholesterol)
	a.Creatinine = merge(a.Creatinine, newPositions.Creatinine)
	a.BloodUreaNitrogen = merge(a.BloodUreaNitrogen, newPositions.BloodUreaNitrogen)
	a.WhiteBloodCellCount = merge(a.WhiteBloodCellCount, newPositions.WhiteBloodCellCount)
	a.Sodium = merge(a.Sodium, newPositions.Sodium)
	a.Potassium = merge(a.Potassium, newPositions.Potassium)
	a.AlanineAminotransferase = merge(a.AlanineAminotransferase, newPositions.AlanineAminotransferase)
	a.AspartateAminotransferase = merge(a.AspartateAminotransferase, newPositions.AspartateAminotransferase)
	a.Calcium = merge(a.Calcium, newPositions.Calcium)
	a.Magnesium = merge(a.Magnesium, newPositions.Magnesium)
	a.ThyroidStimulatingHormone = merge(a.ThyroidStimulatingHormone, newPositions.ThyroidStimulatingHormone)
	a.CReactiveProtein = merge(a.CReactiveProtein, newPositions.CReactiveProtein)
}
