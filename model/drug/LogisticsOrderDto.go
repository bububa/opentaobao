package drug

// LogisticsOrderDto 
type LogisticsOrderDto struct {

    // 快递公司编码
    CpCode   string `json:"cp_code,omitempty"`

    // 快递单号
    MailNo   string `json:"mail_no,omitempty"`

    // 电子面单
    Waybill   string `json:"waybill,omitempty"`

}
