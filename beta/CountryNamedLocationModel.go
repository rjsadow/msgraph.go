// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CountryNamedLocation undocumented
type CountryNamedLocation struct {
	// NamedLocation is the base model of CountryNamedLocation
	NamedLocation
	// CountriesAndRegions undocumented
	CountriesAndRegions []string `json:"countriesAndRegions,omitempty"`
	// IncludeUnknownCountriesAndRegions undocumented
	IncludeUnknownCountriesAndRegions *bool `json:"includeUnknownCountriesAndRegions,omitempty"`
}
