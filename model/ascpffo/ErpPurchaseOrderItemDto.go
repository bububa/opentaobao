package ascpffo

// ErpPurchaseOrderItemDto 
type ErpPurchaseOrderItemDto struct {

    // 扩展字段
    
    ExtendFields   string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
    

    // 未税采购金额，单位元
    
    NoTaxPurchaseAmountDec   string `json:"no_tax_purchase_amount_dec,omitempty" xml:"no_tax_purchase_amount_dec,omitempty"`
    

    // 含税采购金额，单位元
    
    PurchaseAmountDec   string `json:"purchase_amount_dec,omitempty" xml:"purchase_amount_dec,omitempty"`
    

    // 未税采购单价，单位元
    
    NoTaxPurchasePriceDec   string `json:"no_tax_purchase_price_dec,omitempty" xml:"no_tax_purchase_price_dec,omitempty"`
    

    // 含税采购单价，单位元
    
    PurchasePriceDec   string `json:"purchase_price_dec,omitempty" xml:"purchase_price_dec,omitempty"`
    

    // 税率
    
    TaxRate   string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
    

    // 实收残品数量
    
    ReceivedDefectiveQty   int64 `json:"received_defective_qty,omitempty" xml:"received_defective_qty,omitempty"`
    

    // 实收正品数量
    
    ReceivedNormalQty   int64 `json:"received_normal_qty,omitempty" xml:"received_normal_qty,omitempty"`
    

    // 采购数量
    
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    

    // 货品名称
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 货品Id
    
    ScItemId   int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
    

    // 采购单号
    
    PurchaseOrderNo   string `json:"purchase_order_no,omitempty" xml:"purchase_order_no,omitempty"`
    

}
