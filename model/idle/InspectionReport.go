package idle

// InspectionReport 结构体
type InspectionReport struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 鉴定报告
	Report string `json:"report,omitempty" xml:"report,omitempty"`
	// 错误描述
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 成色
	Degree string `json:"degree,omitempty" xml:"degree,omitempty"`
	// 估价id,与order_id任选其一，orderId优先
	QuoteId string `json:"quote_id,omitempty" xml:"quote_id,omitempty"`
	// 对此商品的质检描述
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 特别说明，Json串,一个字符串存储特别说明的文字描述,一个字符串列表存储上传的图片url:json形式:
	Explanation string `json:"explanation,omitempty" xml:"explanation,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 订单，orderId优先，与quote_id任选其一
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
