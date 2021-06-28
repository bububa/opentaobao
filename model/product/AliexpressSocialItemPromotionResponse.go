package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取推广链接 APIResponse
aliexpress.social.item.promotion

获取商品社交推广链接
*/
type AliexpressSocialItemPromotionAPIResponse struct {
    model.CommonResponse
    // Response *AliexpressSocialItemPromotionResponse `json:"aliexpress_social_item_promotion_response,omitempty"` 
    AliexpressSocialItemPromotionResponse
}

/* model for simplify = false
type AliexpressSocialItemPromotionResponse struct {

    // 推广链接
    
    PromotionUrl   string `json:"promotion_url,omitempty"`
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 错误信息
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

}
*/

type AliexpressSocialItemPromotionResponse struct {

    // 推广链接
    PromotionUrl   string `json:"promotion_url,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // 错误信息
    ResultMsg   string `json:"result_msg,omitempty"`

}
