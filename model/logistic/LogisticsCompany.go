package logistic

// LogisticsCompany 
type LogisticsCompany struct {

    // 物流公司标识
    Id   int64 `json:"id,omitempty"`

    // 物流公司代码
    Code   string `json:"code,omitempty"`

    // 物流公司简称
    Name   string `json:"name,omitempty"`

    // 运单号验证正则表达式
    RegMailNo   string `json:"reg_mail_no,omitempty"`

}
