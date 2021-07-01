package dt

// AlibabaDtTmllcarLeadsinfoResults 结构体
type AlibabaDtTmllcarLeadsinfoResults struct {
	// result
	Results []AlibabaDtTmllcarLeadsinfoResult `json:"results,omitempty" xml:"results>alibaba_dt_tmllcar_leadsinfo_result,omitempty"`
	// successed
	Successed bool `json:"successed,omitempty" xml:"successed,omitempty"`
}
