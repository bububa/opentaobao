package wdk

// PromotionContentResult 
/* model for simplify = false
type PromotionContentResult struct {

    // 错误信息
    
    ErrorMessageList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"error_message_list,omitempty"`
    

    // 唯一标示
    
    SkuCodeUniqueIdStr   string `json:"sku_code_unique_id_str,omitempty"`
    

}
*/

// PromotionContentResult 
type PromotionContentResult struct {

    // 错误信息
    ErrorMessageList   []string `json:"error_message_list,omitempty"`

    // 唯一标示
    SkuCodeUniqueIdStr   string `json:"sku_code_unique_id_str,omitempty"`

}
