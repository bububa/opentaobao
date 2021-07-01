package ascpchannel

// AlibabaAscpUopSupplierWaybillQueryData 结构体
type AlibabaAscpUopSupplierWaybillQueryData struct {
	// 请求唯一id
	InvokeId string `json:"invoke_id,omitempty" xml:"invoke_id,omitempty"`
	// 面单号
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
	// 母面单号
	ParentWaybillCode string `json:"parent_waybill_code,omitempty" xml:"parent_waybill_code,omitempty"`
	// 加密面单信息(json字符串)
	PrintData string `json:"print_data,omitempty" xml:"print_data,omitempty"`
	// 模版url
	TemplateURL string `json:"template_u_r_l,omitempty" xml:"template_u_r_l,omitempty"`
	// 签名
	Signature string `json:"signature,omitempty" xml:"signature,omitempty"`
	// 物流单号
	LgOrderCode string `json:"lg_order_code,omitempty" xml:"lg_order_code,omitempty"`
	// 包裹号
	PackageId string `json:"package_id,omitempty" xml:"package_id,omitempty"`
}
