package qimen

// Invoice 
type Invoice struct {

    // 发票抬头
    Header   string `json:"header,omitempty"`

    // 发票金额
    Amount   string `json:"amount,omitempty"`

    // 发票内容
    Content   string `json:"content,omitempty"`

    // 订单详情
    Detail   *Detail `json:"detail,omitempty"`

    // 发票代码(纳税企业的标识)
    Code   string `json:"code,omitempty"`

    // 发票号码(纳税企业内部的发票号)
    Number   string `json:"number,omitempty"`

    // 发票类型(INVOICE=普通发票;VINVOICE=增值税普通发票;EVINVOICE=电子增票;填写的 条件 是:invoiceFlag为Y)
    Type   string `json:"type,omitempty"`

    // 税号
    TaxNumber   string `json:"taxNumber,omitempty"`

}
