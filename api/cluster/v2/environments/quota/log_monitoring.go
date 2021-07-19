package quota

// LogMonitoring represents log monitoring consumption and quota information on environment level. Not set (and not editable) if Log monitoring is not enabled. Not set (and not editable) if Log monitoring is migrated to Davis data on license level. If skipped when editing via PUT method then already set quotas will remain
type LogMonitoring struct {
	MonthlyLimit *int64 `json:"monthlyLimit,omitempty"` // Monthly environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited
	AnnualLimit  *int64 `json:"annualLimit,omitempty"`  // Annual environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited
	// READ ONLY
	ConsumedThisMonth *float64 `json:"consumedThisMonth,omitempty"` // Monthly environment consumption. Resets each calendar month
	ConsumedThisYear  *float64 `json:"consumedThisYear,omitempty"`  // Yearly environment consumption. Resets each year on license creation date anniversary
}
