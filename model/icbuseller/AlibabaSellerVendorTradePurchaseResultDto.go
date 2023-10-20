package icbuseller

// AlibabasellervendortradepurchaseResultDto 结构体
type AlibabasellervendortradepurchaseResultDto struct {
	// 授权订单集合
	Dtos []TradePurchaseDto `json:"dtos,omitempty" xml:"dtos>trade_purchase_dto,omitempty"`
	// 描述
	ExecDescription string `json:"exec_description,omitempty" xml:"exec_description,omitempty"`
	// 接口状态
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
