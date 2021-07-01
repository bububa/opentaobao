package mos

// MultiResult 结构体
type MultiResult struct {
	// total
	ResultTotal int64 `json:"result_total,omitempty" xml:"result_total,omitempty"`
	// errMessage
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// data
	ResultDatas []string `json:"result_datas,omitempty" xml:"result_datas>string,omitempty"`
	// errCode
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// success
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
