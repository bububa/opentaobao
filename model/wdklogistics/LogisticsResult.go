package wdklogistics

// LogisticsResult 结构体
type LogisticsResult struct {
	// 内容节点
	Datas []AlibabawdklogisticspuspickupcararrivedData `json:"datas,omitempty" xml:"datas>alibabawdklogisticspuspickupcararrived_data,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 接口服务查询结果描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
