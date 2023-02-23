package cloudgame

// AlibabaCloudgameOpenidQueryResult 结构体
type AlibabaCloudgameOpenidQueryResult struct {
	// 结果码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果码描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回数据
	Data *HavanaUserIdQueryResponseVo `json:"data,omitempty" xml:"data,omitempty"`
}
