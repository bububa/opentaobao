package wdk

// PromotionContentResult 结构体
type PromotionContentResult struct {
	// 错误信息
	ErrorMessageList []string `json:"error_message_list,omitempty" xml:"error_message_list>string,omitempty"`
	// 唯一标示
	SkuCodeUniqueIdStr string `json:"sku_code_unique_id_str,omitempty" xml:"sku_code_unique_id_str,omitempty"`
}
