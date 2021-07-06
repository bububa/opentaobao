package icbulogistics

// AlibabaOnetouchLogisticsExpressAddressCityListResult 结构体
type AlibabaOnetouchLogisticsExpressAddressCityListResult struct {
	// 列表对象
	Values []RegionEntity `json:"values,omitempty" xml:"values>region_entity,omitempty"`
	// 返回结果描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 返回结果编码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
