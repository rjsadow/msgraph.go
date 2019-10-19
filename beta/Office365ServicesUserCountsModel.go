// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// Office365ServicesUserCounts undocumented
type Office365ServicesUserCounts struct {
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *time.Time `json:"reportRefreshDate,omitempty"`
	// ExchangeActive undocumented
	ExchangeActive *int `json:"exchangeActive,omitempty"`
	// ExchangeInactive undocumented
	ExchangeInactive *int `json:"exchangeInactive,omitempty"`
	// OneDriveActive undocumented
	OneDriveActive *int `json:"oneDriveActive,omitempty"`
	// OneDriveInactive undocumented
	OneDriveInactive *int `json:"oneDriveInactive,omitempty"`
	// SharePointActive undocumented
	SharePointActive *int `json:"sharePointActive,omitempty"`
	// SharePointInactive undocumented
	SharePointInactive *int `json:"sharePointInactive,omitempty"`
	// SkypeForBusinessActive undocumented
	SkypeForBusinessActive *int `json:"skypeForBusinessActive,omitempty"`
	// SkypeForBusinessInactive undocumented
	SkypeForBusinessInactive *int `json:"skypeForBusinessInactive,omitempty"`
	// YammerActive undocumented
	YammerActive *int `json:"yammerActive,omitempty"`
	// YammerInactive undocumented
	YammerInactive *int `json:"yammerInactive,omitempty"`
	// TeamsActive undocumented
	TeamsActive *int `json:"teamsActive,omitempty"`
	// TeamsInactive undocumented
	TeamsInactive *int `json:"teamsInactive,omitempty"`
	// Office365Active undocumented
	Office365Active *int `json:"office365Active,omitempty"`
	// Office365Inactive undocumented
	Office365Inactive *int `json:"office365Inactive,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}