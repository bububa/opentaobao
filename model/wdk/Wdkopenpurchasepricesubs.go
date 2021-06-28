package wdk

// Wdkopenpurchasepricesubs 
type Wdkopenpurchasepricesubs struct {

    // 确认标识，0:核对 1:确认，必填
    
    Confirm   string `json:"confirm,omitempty" xml:"confirm,omitempty"`
    

    // 去税采购价，单位分，必填
    
    PriceWithoutTax   string `json:"price_without_tax,omitempty" xml:"price_without_tax,omitempty"`
    

    // 含税采购价，单位分，必填
    
    PriceWithTax   string `json:"price_with_tax,omitempty" xml:"price_with_tax,omitempty"`
    

    // 税率，必填
    
    TaxRate   string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
    

    // 商品编码，必填
    
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    

    // 淘系子订单号，必填
    
    TbSubOrderId   string `json:"tb_sub_order_id,omitempty" xml:"tb_sub_order_id,omitempty"`
    

}
