package tmallgenie

// OpenInfoResponse 结构体
type OpenInfoResponse struct {
	// 关联id
	UnionIds []UnionIdInfo `json:"union_ids,omitempty" xml:"union_ids>union_id_info,omitempty"`
	// 开放openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 入参内容
	Param *ConverterIdRequest `json:"param,omitempty" xml:"param,omitempty"`
}
