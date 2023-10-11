package simba

// TopServiceContext 结构体
type TopServiceContext struct {
	// api业务线编码,查询账户余额bizCode必须是universalBP
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
}
