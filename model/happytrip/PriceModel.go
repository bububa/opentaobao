package happytrip

// PriceModel 
type PriceModel struct {

    // 起步价格 单位：元
    
    StartPrice   string `json:"start_price,omitempty" xml:"start_price,omitempty"`
    

    // 每公里单价 单位：元
    
    NormalUnitPrice   string `json:"normal_unit_price,omitempty" xml:"normal_unit_price,omitempty"`
    

    // 动调溢价 单位：元
    
    DynamicPrice   string `json:"dynamic_price,omitempty" xml:"dynamic_price,omitempty"`
    

    // 价格提示
    
    PriceTip   string `json:"price_tip,omitempty" xml:"price_tip,omitempty"`
    

    // 总价格(包含dynamic_price) 单位：元，如果有折扣，这里为折扣后价格，如果没有折扣这里和original_price保持一致；允许拼车时为拼成一口价
    
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    

    // 动态调价md5，用于锁定订单创建时的价格
    
    DynamicMd5   string `json:"dynamic_md5,omitempty" xml:"dynamic_md5,omitempty"`
    

    // 供应商的车型名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 供应商的车型代码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 原始价格，折扣优惠前的价格；允许拼车时为未拼成一口价
    
    OriginalPrice   string `json:"original_price,omitempty" xml:"original_price,omitempty"`
    

    // 线路类型，为空或0，表示普通线路；1，表示专线
    
    LineType   int64 `json:"line_type,omitempty" xml:"line_type,omitempty"`
    

    // 线路信息，线路为专线时不为空
    
    LineInfo   *LineInfo `json:"line_info,omitempty" xml:"line_info,omitempty"`
    

}
