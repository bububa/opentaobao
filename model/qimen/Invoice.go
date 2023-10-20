package qimen

// Invoice 结构体
type Invoice struct {
	// 发票抬头
	Header string `json:"header,omitempty" xml:"header,omitempty"`
	// 发票金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 发票内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 发票代码(纳税企业的标识)
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 发票号码(纳税企业内部的发票号)
	Number string `json:"number,omitempty" xml:"number,omitempty"`
	// 发票类型(INVOICE=普通发票;VINVOICE=增值税普通发票;EVINVOICE=电子增票;填写的 条件 是:invoiceFlag为Y)
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 税号
	TaxNumber string `json:"taxNumber,omitempty" xml:"taxNumber,omitempty"`
	// 订单详情
	Detail *TaobaoqimendeliveryorderbatchconfirmDetail `json:"detail,omitempty" xml:"detail,omitempty"`
}
