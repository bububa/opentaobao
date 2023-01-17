package tblogistics

// TopConsignPkgRequest 结构体
type TopConsignPkgRequest struct {
	// 物流公司代码.如&#34;POST&#34;代表中国邮政,&#34;ZJS&#34;代表宅急送。调用 taobao.logistics.companies.get 获取
	CompanyCode string `json:"company_code,omitempty" xml:"company_code,omitempty"`
	// 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入
	OutSid string `json:"out_sid,omitempty" xml:"out_sid,omitempty"`
}
