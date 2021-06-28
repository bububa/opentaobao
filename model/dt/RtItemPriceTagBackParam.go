package dt

// RtItemPriceTagBackParam 
/* model for simplify = false
type RtItemPriceTagBackParam struct {

    // 数据列表
    
    PriceTagParamList  struct {
        RtItemResearchPriceParam  []RtItemResearchPriceParam `json:"rt_item_research_price_param,omitempty"`
    } `json:"price_tag_param_list,omitempty"`
    

    // 业务编码
    
    BusiCode   string `json:"busi_code,omitempty"`
    

    // 业务来源
    
    Source   string `json:"source,omitempty"`
    

}
*/

// RtItemPriceTagBackParam 
type RtItemPriceTagBackParam struct {

    // 数据列表
    PriceTagParamList   []RtItemResearchPriceParam `json:"price_tag_param_list,omitempty"`

    // 业务编码
    BusiCode   string `json:"busi_code,omitempty"`

    // 业务来源
    Source   string `json:"source,omitempty"`

}
