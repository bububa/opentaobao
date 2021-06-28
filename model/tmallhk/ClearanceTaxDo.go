package tmallhk

// ClearanceTaxDo 
/* model for simplify = false
type ClearanceTaxDo struct {

    // 关税，主&子
    
    CustomDutyFee   int64 `json:"custom_duty_fee,omitempty"`
    

    // 优惠，主&子
    
    CustomsCouponFee   int64 `json:"customs_coupon_fee,omitempty"`
    

    // 保费，主&子
    
    CustomsInsuranceFee   int64 `json:"customs_insurance_fee,omitempty"`
    

    // 完税价，子订单
    
    CustomsSubTotalFee   int64 `json:"customs_sub_total_fee,omitempty"`
    

    // 完税价，主订单
    
    CustomsTotalFee   int64 `json:"customs_total_fee,omitempty"`
    

    // 消费税，主&子
    
    ExciseDutyFee   int64 `json:"excise_duty_fee,omitempty"`
    

    // 海关税收编码
    
    Hscode   string `json:"hscode,omitempty"`
    

    // 税费，子订单
    
    OrderLineTotalTaxFee   int64 `json:"order_line_total_tax_fee,omitempty"`
    

    // 总税费，主订单
    
    OrderTotalTaxFee   int64 `json:"order_total_tax_fee,omitempty"`
    

    // 邮费，主&子
    
    PostFee   int64 `json:"post_fee,omitempty"`
    

    // 给海关的税费中的关税，主&子
    
    TariffCustomFee   int64 `json:"tariff_custom_fee,omitempty"`
    

    // 给海关的税费中的消费税，主&子
    
    TariffExciseFee   int64 `json:"tariff_excise_fee,omitempty"`
    

    // 给海关的税费，主&子
    
    TariffFee   int64 `json:"tariff_fee,omitempty"`
    

    // 给海关的税费中的增值税，主&子
    
    TariffVatFee   int64 `json:"tariff_vat_fee,omitempty"`
    

    // 增值税，主&子
    
    VatFee   int64 `json:"vat_fee,omitempty"`
    

}
*/

// ClearanceTaxDo 
type ClearanceTaxDo struct {

    // 关税，主&子
    CustomDutyFee   int64 `json:"custom_duty_fee,omitempty"`

    // 优惠，主&子
    CustomsCouponFee   int64 `json:"customs_coupon_fee,omitempty"`

    // 保费，主&子
    CustomsInsuranceFee   int64 `json:"customs_insurance_fee,omitempty"`

    // 完税价，子订单
    CustomsSubTotalFee   int64 `json:"customs_sub_total_fee,omitempty"`

    // 完税价，主订单
    CustomsTotalFee   int64 `json:"customs_total_fee,omitempty"`

    // 消费税，主&子
    ExciseDutyFee   int64 `json:"excise_duty_fee,omitempty"`

    // 海关税收编码
    Hscode   string `json:"hscode,omitempty"`

    // 税费，子订单
    OrderLineTotalTaxFee   int64 `json:"order_line_total_tax_fee,omitempty"`

    // 总税费，主订单
    OrderTotalTaxFee   int64 `json:"order_total_tax_fee,omitempty"`

    // 邮费，主&子
    PostFee   int64 `json:"post_fee,omitempty"`

    // 给海关的税费中的关税，主&子
    TariffCustomFee   int64 `json:"tariff_custom_fee,omitempty"`

    // 给海关的税费中的消费税，主&子
    TariffExciseFee   int64 `json:"tariff_excise_fee,omitempty"`

    // 给海关的税费，主&子
    TariffFee   int64 `json:"tariff_fee,omitempty"`

    // 给海关的税费中的增值税，主&子
    TariffVatFee   int64 `json:"tariff_vat_fee,omitempty"`

    // 增值税，主&子
    VatFee   int64 `json:"vat_fee,omitempty"`

}
