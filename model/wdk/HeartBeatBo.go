package wdk

// HeartBeatBo 结构体
type HeartBeatBo struct {
	// MARKET-营销，ITEM-商品
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 当前版本信息
	VersionId int64 `json:"version_id,omitempty" xml:"version_id,omitempty"`
}
