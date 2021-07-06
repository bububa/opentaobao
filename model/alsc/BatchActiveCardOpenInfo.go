package alsc

// BatchActiveCardOpenInfo 结构体
type BatchActiveCardOpenInfo struct {
	// 结果 < KEY：Id  VALUE：描述（SUCCESS-通过） >
	ResultMap string `json:"result_map,omitempty" xml:"result_map,omitempty"`
	// 是否全部激活成功
	AllSuccess bool `json:"all_success,omitempty" xml:"all_success,omitempty"`
}
