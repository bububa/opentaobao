package mos

// BillSettlementDto 
type BillSettlementDto struct {

    // 行号
    LineNo   string `json:"line_no,omitempty"`

    // 结算单金额
    Amount   string `json:"amount,omitempty"`

    // 结算行税率，普票的税率必须是0，专票的税率不能为0
    TaxRate   string `json:"tax_rate,omitempty"`

    // 业务细类代码
    BizSubModuleCode   string `json:"biz_sub_module_code,omitempty"`

    // 发票类型 SPECIAL("专票")，ORDINARY("普票")
    InvoiceType   string `json:"invoice_type,omitempty"`

    // 供应商id（可空）
    SupplierNo   string `json:"supplier_no,omitempty"`

    // 供应商名称（可空
    SupplierName   string `json:"supplier_name,omitempty"`

    // 收款方开户省
    BankProvince   string `json:"bank_province,omitempty"`

    // 收款方开户市
    BankCity   string `json:"bank_city,omitempty"`

    // 收款方银行账号
    AccountNo   string `json:"account_no,omitempty"`

    // 收款方账号名称
    AccountName   string `json:"account_name,omitempty"`

    // 银行号
    BankCode   string `json:"bank_code,omitempty"`

    // 收款方开户行
    BankName   string `json:"bank_name,omitempty"`

    // 收款方开户支行
    BankBranchName   string `json:"bank_branch_name,omitempty"`

    // 收款方开户支行code
    BankBranchCode   string `json:"bank_branch_code,omitempty"`

    // 账号类型，COMPANY企业/PERSON个人
    AccountTypes   string `json:"account_types,omitempty"`

    // 联行号
    CnapsCode   string `json:"cnaps_code,omitempty"`

    // 联系人
    Contactor   string `json:"contactor,omitempty"`

    // 联系方式：国际化区号 + 座机/手机
    Telephone   string `json:"telephone,omitempty"`

    // 付款说明。该字段会通过银行传给供应商，过长时会自动截取
    Comments   string `json:"comments,omitempty"`

    // 扩展
    ExtendParams   string `json:"extend_params,omitempty"`

    // 发票列表
    InvoiceDTOList   []SettleInvoiceDto `json:"invoice_d_t_o_list,omitempty"`

    // 受益部门coa
    DepartmentCoa   string `json:"department_coa,omitempty"`

    // 区域科目段
    CityCoa   string `json:"city_coa,omitempty"`

}
