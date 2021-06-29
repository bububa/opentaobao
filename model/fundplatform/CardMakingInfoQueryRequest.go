package fundplatform

// CardMakingInfoQueryRequest 
type CardMakingInfoQueryRequest struct {

    // 页大小,最大500
    
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    

    // 当前页，从1开始
    
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    

    // 已废弃,可以填写固定值test
    
    Signture   string `json:"signture,omitempty" xml:"signture,omitempty"`
    

    // 子制卡单ID
    
    CardOrderId   int64 `json:"card_order_id,omitempty" xml:"card_order_id,omitempty"`
    

    // 卡号列表
    
    CardNos   []string `json:"card_nos,omitempty" xml:"card_nos>string,omitempty"`
    

}
