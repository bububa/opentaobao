package scs

// ApiServiceContext 结构体
type ApiServiceContext struct {
	// api业务线编码。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
}
