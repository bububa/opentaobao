package icbuassurance

// AssuranceFlag 结构体
type AssuranceFlag struct {
	// flagList
	FlagList []string `json:"flag_list,omitempty" xml:"flag_list>string,omitempty"`
	// guideURL
	GuideURL string `json:"guide_u_r_l,omitempty" xml:"guide_u_r_l,omitempty"`
	// pauseStatus
	PauseStatus bool `json:"pause_status,omitempty" xml:"pause_status,omitempty"`
}
