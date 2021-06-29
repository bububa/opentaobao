package einvoice

// UserInvoiceApplyDTO 
type UserInvoiceApplyDTO struct {
    // 开票申请id
    ApplyId   string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    // 开票金额
    InvoiceAmount   string `json:"invoice_amount,omitempty" xml:"invoice_amount,omitempty"`
    // 税号，表示商家是为此税号的账单开票
    PayeeRegisterNo   string `json:"payee_register_no,omitempty" xml:"payee_register_no,omitempty"`
    // 企业名称
    CompanyName   string `json:"company_name,omitempty" xml:"company_name,omitempty"`
    // 申请企业开户行
    Bank   string `json:"bank,omitempty" xml:"bank,omitempty"`
    // 申请企业开户账号
    BankAccount   string `json:"bank_account,omitempty" xml:"bank_account,omitempty"`
    // 申请单状态：1：待确认，2：开票中，3：拒绝开票，4：发票已发出，0：完成开票
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 商家收货地址
    ReceiverAddress   string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
    // 商家收货人
    ReceiverName   string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
    // 商家收货电话
    ReceiverPhone   string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
    // 服务商发货人
    SenderName   string `json:"sender_name,omitempty" xml:"sender_name,omitempty"`
    // 服务商发货人电话
    SenderPhone   string `json:"sender_phone,omitempty" xml:"sender_phone,omitempty"`
    // 服务商发货物流
    SenderLogisticsCompany   string `json:"sender_logistics_company,omitempty" xml:"sender_logistics_company,omitempty"`
    // 服务商发货快递单号
    SenderLogisticsNo   string `json:"sender_logistics_no,omitempty" xml:"sender_logistics_no,omitempty"`
    // 发票类型：1:增值税普通发票，2:增值税专用发票
    InvoiceType   int64 `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
    // 开票明细列表
    InvoiceItemList   []UserInvoiceItemDTO `json:"invoice_item_list,omitempty" xml:"invoice_item_list>user_invoice_item_dto,omitempty"`
    // 购方票面税号
    InvoicePayeeRegisterNo   string `json:"invoice_payee_register_no,omitempty" xml:"invoice_payee_register_no,omitempty"`
    // 购方票面地址
    InvoiceAddress   string `json:"invoice_address,omitempty" xml:"invoice_address,omitempty"`
    // 购方票面电话
    InvoicePhone   string `json:"invoice_phone,omitempty" xml:"invoice_phone,omitempty"`
}
