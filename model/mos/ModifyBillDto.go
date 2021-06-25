package mos

// ModifyBillDto 
type ModifyBillDto struct {

    // 收款方开户支行code
    BankBranchCode   string `json:"bank_branch_code,omitempty"`

    // 供应商名称（可空）
    SupplierName   string `json:"supplier_name,omitempty"`

    // 供应商id（可空
    SupplierNo   string `json:"supplier_no,omitempty"`

    // 收款方开户市
    BankCity   string `json:"bank_city,omitempty"`

    // 收款方银行账号
    AccountNo   string `json:"account_no,omitempty"`

    // 收款方开户行
    BankName   string `json:"bank_name,omitempty"`

    // 账号类型，COMPANY企业/PERSON个人
    AccountTypes   string `json:"account_types,omitempty"`

    // 收款方开户省
    BankProvince   string `json:"bank_province,omitempty"`

    // //银行号
    BankCode   string `json:"bank_code,omitempty"`

    // 结算行号list
    SettleLineNos   []String `json:"settle_line_nos,omitempty"`

    // 收款方账号名称
    AccountName   string `json:"account_name,omitempty"`

    // 收款方开户支行
    BankBranchName   string `json:"bank_branch_name,omitempty"`

    // 联行号
    CnapsCode   string `json:"cnaps_code,omitempty"`

    // //单据请求唯一编号，必填
    BillNo   string `json:"bill_no,omitempty"`

    // 扩展参数
    ExtendParams   string `json:"extend_params,omitempty"`

}
