package bus

// CitySearchRp 结构体
type CitySearchRp struct {
	// 城市集合
	Citys []CityDto `json:"citys,omitempty" xml:"citys>city_dto,omitempty"`
	// 错误代码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
