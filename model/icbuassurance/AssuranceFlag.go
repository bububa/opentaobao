package icbuassurance

// AssuranceFlag 结构体
type AssuranceFlag struct {
	// pauseStatus
	PauseStatus bool `json:"pause_status,omitempty" xml:"pause_status,omitempty"`
	// guideURL
	GuideURL string `json:"guide_u_r_l,omitempty" xml:"guide_u_r_l,omitempty"`
	// flagList
	FlagList []string `json:"flag_list,omitempty" xml:"flag_list>string,omitempty"`
}
