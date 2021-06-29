package seaking

// Extra 
type Extra struct {
    // 店铺所在平台
    SubIdentifier   string `json:"sub_identifier,omitempty" xml:"sub_identifier,omitempty"`
    // 店铺id
    SubIdentifyType   string `json:"sub_identify_type,omitempty" xml:"sub_identify_type,omitempty"`
    // 商品id
    ProductId   string `json:"product_id,omitempty" xml:"product_id,omitempty"`
    // 商品所在平台
    Platform   string `json:"platform,omitempty" xml:"platform,omitempty"`
    // 是否需要orc信息
    HaveOrc   string `json:"have_orc,omitempty" xml:"have_orc,omitempty"`
}
