package alsc

// QueryPhyCardOpenReq 
type QueryPhyCardOpenReq struct {
    // 外部品牌ID
    OutBrandId   string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
    // 物理卡号
    PhysicalCardId   string `json:"physical_card_id,omitempty" xml:"physical_card_id,omitempty"`
    // 品牌ID
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
}
