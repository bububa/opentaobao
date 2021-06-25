package bus

// CitySearchRp 
type CitySearchRp struct {

    // 城市集合
    Citys   []CityDto `json:"citys,omitempty"`

    // 错误代码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误描述
    ErrMsg   string `json:"err_msg,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
