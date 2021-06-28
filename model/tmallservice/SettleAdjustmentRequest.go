package tmallservice

// SettleAdjustmentRequest 
type SettleAdjustmentRequest struct {

    // 调整费用，必需是正数，单位分
    
    Cost   int64 `json:"cost,omitempty" xml:"cost,omitempty"`
    

    // 调整原因描述
    
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    

    // 调整原因图片url,最后不用加分号，最多三条
    
    PictureUrls   string `json:"picture_urls,omitempty" xml:"picture_urls,omitempty"`
    

    // 计价因子，填写规则：1、有计价因子场景：{name:计价因子名称 ,value:数量｝如示例；2、没有计价因子场景：填默认值：｛name:计价因子,value:0｝
    
    PriceFactors   []SettlementPriceFactor `json:"price_factors,omitempty" xml:"price_factors,omitempty"`
    

    // 调整单分类类型；1,配件费;2,不符单费;3,拆旧费;4,二次上门;5,胶费;6,打孔费;7,层高费;8,远程费;9,单外费;10,其他
    
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    

    // 工单ID
    
    WorkcardId   int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
    

}
