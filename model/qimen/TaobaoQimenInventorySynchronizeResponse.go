package qimen

// TaobaoQimenInventorySynchronizeResponse 结构体
type TaobaoQimenInventorySynchronizeResponse struct {
	// 响应结果,success|failure,string (10),必填
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码,,string (50),
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息,,string (100),
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
