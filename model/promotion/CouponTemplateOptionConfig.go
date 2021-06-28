package promotion

// CouponTemplateOptionConfig 
type CouponTemplateOptionConfig struct {

    // 外部优惠券模板id
    
    OutCouponTemplateId   string `json:"out_coupon_template_id,omitempty" xml:"out_coupon_template_id,omitempty"`
    

    // 外部优惠券类型
    
    OutCouponType   string `json:"out_coupon_type,omitempty" xml:"out_coupon_type,omitempty"`
    

    // 优惠券logo url
    
    LogoUrl   string `json:"logo_url,omitempty" xml:"logo_url,omitempty"`
    

    // 券图片
    
    PictureUrl   string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
    

    // 合同ID
    
    ContractInstanceId   int64 `json:"contract_instance_id,omitempty" xml:"contract_instance_id,omitempty"`
    

    // 采买者ID
    
    PurchaseId   int64 `json:"purchase_id,omitempty" xml:"purchase_id,omitempty"`
    

    // 抵用券是否可贬值使用
    
    Devalue   bool `json:"devalue,omitempty" xml:"devalue,omitempty"`
    

}
