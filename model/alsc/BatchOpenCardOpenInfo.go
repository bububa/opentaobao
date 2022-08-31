package alsc

// BatchOpenCardOpenInfo 结构体
type BatchOpenCardOpenInfo struct {
	// 结果 &lt; KEY：Id  VALUE：描述（SUCCESS-通过） &gt;
	ResultMap string `json:"result_map,omitempty" xml:"result_map,omitempty"`
	// 是否全部开通成功
	AllSuccess string `json:"all_success,omitempty" xml:"all_success,omitempty"`
}
