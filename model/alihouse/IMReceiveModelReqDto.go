package alihouse

// IMReceiveModelReqDto 结构体
type IMReceiveModelReqDto struct {
	// 外部id列表，根据type区分
	TagIds []string `json:"tag_ids,omitempty" xml:"tag_ids>string,omitempty"`
	// 0-楼盘 1-交易商品
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 1-天猫好房来客通 2-千牛
	ImModel int64 `json:"im_model,omitempty" xml:"im_model,omitempty"`
	// 0-非测试 1-测试
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}
