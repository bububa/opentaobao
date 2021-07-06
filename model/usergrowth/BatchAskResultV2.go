package usergrowth

// BatchAskResultV2 结构体
type BatchAskResultV2 struct {
	// 匹配的设备与其任务信息列表
	Results []BatchAskResultItem `json:"results,omitempty" xml:"results>batch_ask_result_item,omitempty"`
	// 错误码， 0： 成功；1：限流；2：服务不可用
	Errcode int64 `json:"errcode,omitempty" xml:"errcode,omitempty"`
}
