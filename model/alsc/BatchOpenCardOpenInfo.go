package alsc

// BatchOpenCardOpenInfo 结构体
type BatchOpenCardOpenInfo struct {
	// 是否全部开通成功
	AllSuccess string `json:"all_success,omitempty" xml:"all_success,omitempty"`
	// 结果 < KEY：Id  VALUE：描述（SUCCESS-通过） >
	ResultMap string `json:"result_map,omitempty" xml:"result_map,omitempty"`
}
