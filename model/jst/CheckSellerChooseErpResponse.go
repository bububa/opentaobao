package jst

// CheckSellerChooseErpResponse 结构体
type CheckSellerChooseErpResponse struct {
	// 状态结果&lt;/br&gt; 您无权查询此商家信息&lt;br&gt; 已与当前服务商签约改地址服务&lt;/br&gt;  当前商家改地址服务未签约服务商&lt;/br&gt; 已与其他服务商签约改地址服务&lt;/br&gt;
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 商家绑定的appkey
	Appkey string `json:"appkey,omitempty" xml:"appkey,omitempty"`
	// app名称
	AppTitle string `json:"app_title,omitempty" xml:"app_title,omitempty"`
}
