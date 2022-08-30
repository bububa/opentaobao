package ascp

// QueryScItemRequest 结构体
type QueryScItemRequest struct {
	// 货品商家编码列表，一次查询不要超过100个货品
	ScitemCodes []string `json:"scitem_codes,omitempty" xml:"scitem_codes>string,omitempty"`
	// 业务请求ID，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务请求时间戳
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
