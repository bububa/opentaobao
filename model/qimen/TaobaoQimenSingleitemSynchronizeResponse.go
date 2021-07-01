package qimen

// TaobaoQimenSingleitemSynchronizeResponse 结构体
type TaobaoQimenSingleitemSynchronizeResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 仓储系统商品Id(当这个字段不为空的时候;所有erp传输的时候都碰到itemid必传)
	ItemId string `json:"itemId,omitempty" xml:"itemId,omitempty"`
}
