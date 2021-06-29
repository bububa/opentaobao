package pur

// SupplierPreInvoiceInfoVO 
type SupplierPreInvoiceInfoVO struct {

    // 预发票行信息
    
    InvoiceItemList   []SupplierPreInvoiceItem `json:"invoice_item_list,omitempty" xml:"invoice_item_list,omitempty"`
    

    // 应付发票状态
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 应付发票备注
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 开票人
    
    Drawer   string `json:"drawer,omitempty" xml:"drawer,omitempty"`
    

    // 复核人
    
    Reviewer   string `json:"reviewer,omitempty" xml:"reviewer,omitempty"`
    

    // 收款人
    
    Payee   string `json:"payee,omitempty" xml:"payee,omitempty"`
    

    // 币种
    
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
    

    // 发票附件
    
    Attachment   []*model.File `json:"attachment,omitempty" xml:"attachment>*model.File,omitempty"`
    

    // 关闭原因
    
    CloseReason   string `json:"close_reason,omitempty" xml:"close_reason,omitempty"`
    

    // 价税合计
    
    TotalAmount   string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
    

    // 税额合计
    
    TaxAmount   string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
    

    // 卖方账户
    
    InvoiceSellerAccountNo   string `json:"invoice_seller_account_no,omitempty" xml:"invoice_seller_account_no,omitempty"`
    

    // 卖方开户行
    
    InvoiceSellerBankNo   string `json:"invoice_seller_bank_no,omitempty" xml:"invoice_seller_bank_no,omitempty"`
    

    // 卖方电话
    
    InvoiceSellerPhone   string `json:"invoice_seller_phone,omitempty" xml:"invoice_seller_phone,omitempty"`
    

    // 卖方地址
    
    InvoiceSellerAddress   string `json:"invoice_seller_address,omitempty" xml:"invoice_seller_address,omitempty"`
    

    // 纳税人识别号
    
    InvoiceSellerTaxCode   string `json:"invoice_seller_tax_code,omitempty" xml:"invoice_seller_tax_code,omitempty"`
    

    // 开票方抬头
    
    InvoiceSellerName   string `json:"invoice_seller_name,omitempty" xml:"invoice_seller_name,omitempty"`
    

    // 购买方账户
    
    InvoiceBuyerAccountNo   string `json:"invoice_buyer_account_no,omitempty" xml:"invoice_buyer_account_no,omitempty"`
    

    // 购买方开户行
    
    InvoiceBuyerBankNo   string `json:"invoice_buyer_bank_no,omitempty" xml:"invoice_buyer_bank_no,omitempty"`
    

    // 购买方电话
    
    InvoiceBuyerPhone   string `json:"invoice_buyer_phone,omitempty" xml:"invoice_buyer_phone,omitempty"`
    

    // 购买方地址
    
    InvoiceBuyerAddress   string `json:"invoice_buyer_address,omitempty" xml:"invoice_buyer_address,omitempty"`
    

    // 纳税人识别号
    
    InvoiceBuyerTaxCode   string `json:"invoice_buyer_tax_code,omitempty" xml:"invoice_buyer_tax_code,omitempty"`
    

    // 开票方抬头
    
    InvoiceBuyerName   string `json:"invoice_buyer_name,omitempty" xml:"invoice_buyer_name,omitempty"`
    

    // 发票类型
    
    InvoiceType   string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
    

    // 发票代码
    
    InvoiceCode   string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
    

    // 发票号码
    
    InvoiceNo   string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
    

    // 开票日期
    
    DayAfterInvoicing   string `json:"day_after_invoicing,omitempty" xml:"day_after_invoicing,omitempty"`
    

    // 采购组织
    
    DemanderPurchaseOrgCode   string `json:"demander_purchase_org_code,omitempty" xml:"demander_purchase_org_code,omitempty"`
    

    // 供应商ou代码
    
    OuCode   string `json:"ou_code,omitempty" xml:"ou_code,omitempty"`
    

    // 供应商名称
    
    SupplierName   string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
    

    // 供应商编码
    
    SupplierCode   string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
    

}
