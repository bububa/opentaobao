package dt

// RtItemPriceTagBackParam 
type RtItemPriceTagBackParam struct {

    // 数据列表
    
    PriceTagParamList   []RtItemResearchPriceParam `json:"price_tag_param_list,omitempty" xml:"price_tag_param_list,omitempty"`
    

    // 业务编码
    
    BusiCode   string `json:"busi_code,omitempty" xml:"busi_code,omitempty"`
    

    // 业务来源
    
    Source   string `json:"source,omitempty" xml:"source,omitempty"`
    

}
