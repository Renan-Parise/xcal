package analytes

import (
	"time"
)

func (a *Analytes) PopulateIncludedAtCreation(t time.Time) {
	setIncludedAt := func(values []Information) {
		for i := range values {
			if values[i].IncludedAt.IsZero() {
				values[i].IncludedAt = t
			}
		}
	}

	setIncludedAt(a.Glucose.Values)
	setIncludedAt(a.Hemoglobin.Values)
	setIncludedAt(a.Cholesterol.Values)
	setIncludedAt(a.Creatinine.Values)
	setIncludedAt(a.BloodUreaNitrogen.Values)
	setIncludedAt(a.WhiteBloodCellCount.Values)
	setIncludedAt(a.Sodium.Values)
	setIncludedAt(a.Potassium.Values)
	setIncludedAt(a.AlanineAminotransferase.Values)
	setIncludedAt(a.AspartateAminotransferase.Values)
	setIncludedAt(a.Calcium.Values)
	setIncludedAt(a.Magnesium.Values)
	setIncludedAt(a.ThyroidStimulatingHormone.Values)
	setIncludedAt(a.CReactiveProtein.Values)
}
