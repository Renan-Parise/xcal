package utils

import (
	"reflect"
	"time"

	"github.com/renan-parise/xcal-analytics/internal/analytes"
)

func PopulateIncludedAtCreation(analytes *analytes.Analytes, t time.Time) {
	setIncludedAt := func(v reflect.Value) {
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			if item.Kind() == reflect.Struct {
				includedAtField := item.FieldByName("IncludedAt")
				if includedAtField.IsValid() && includedAtField.CanSet() {
					includedAtField.Set(reflect.ValueOf(t))
				}
			}
		}
	}

	val := reflect.ValueOf(analytes).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.Struct {
			valuesField := field.FieldByName("Values")
			if valuesField.IsValid() && valuesField.Kind() == reflect.Slice {
				setIncludedAt(valuesField)
			}
		}
	}
}
