package icburfq

// PriceList 
type PriceList struct {

    // 目的港
    
    Port   string `json:"port,omitempty" xml:"port,omitempty"`
    

    // 发运条件
    
    ShippingTerms   string `json:"shipping_terms,omitempty" xml:"shipping_terms,omitempty"`
    

    // 图片image_str,请通过调用alibaba.icbu.annex.upload结果作为入参如果是都个附件通过\r\n分割
    
    ImageStr   string `json:"image_str,omitempty" xml:"image_str,omitempty"`
    

    // 产品编号
    
    ModelNum   string `json:"model_num,omitempty" xml:"model_num,omitempty"`
    

    // 产品名称
    
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    

    // 价格
    
    FobPrice   string `json:"fob_price,omitempty" xml:"fob_price,omitempty"`
    

    // 数量
    
    Quantity   string `json:"quantity,omitempty" xml:"quantity,omitempty"`
    

    // 数量单位
    
    QuantityUnit   string `json:"quantity_unit,omitempty" xml:"quantity_unit,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 价格单位
    
    FobPriceUnit   string `json:"fob_price_unit,omitempty" xml:"fob_price_unit,omitempty"`
    

}
