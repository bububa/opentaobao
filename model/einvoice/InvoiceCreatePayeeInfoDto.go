package einvoice

// InvoiceCreatePayeeInfoDTO 
type InvoiceCreatePayeeInfoDTO struct {
    // 销方地址，销方电话加地址不超出100字符
    PayeeAddress   string `json:"payee_address,omitempty" xml:"payee_address,omitempty"`
    // 销方银行帐号
    PayeeBankAccountId   string `json:"payee_bank_account_id,omitempty" xml:"payee_bank_account_id,omitempty"`
    // 销方开户行名称，开户行账号加名称不超出100字符
    PayeeBankName   string `json:"payee_bank_name,omitempty" xml:"payee_bank_name,omitempty"`
    // 复核人
    PayeeChecker   string `json:"payee_checker,omitempty" xml:"payee_checker,omitempty"`
    // 销方名称，公司名
    PayeeName   string `json:"payee_name,omitempty" xml:"payee_name,omitempty"`
    // 开票人
    PayeeOperator   string `json:"payee_operator,omitempty" xml:"payee_operator,omitempty"`
    // 销方电话
    PayeePhone   string `json:"payee_phone,omitempty" xml:"payee_phone,omitempty"`
    // 收款人
    PayeeReceiver   string `json:"payee_receiver,omitempty" xml:"payee_receiver,omitempty"`
    // 销方税务登记证号
    PayeeRegisterNo   string `json:"payee_register_no,omitempty" xml:"payee_register_no,omitempty"`
}
