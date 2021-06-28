package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推广链接 APIResponse
aliexpress.social.item.promotion

获取商品社交推广链接
*/
type AliexpressSocialItemPromotionAPIResponse struct {
    model.CommonResponse
    AliexpressSocialItemPromotionResponse
}

type AliexpressSocialItemPromotionResponse struct {
    XMLName xml.Name `xml:"aliexpress_social_item_promotion_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 推广链接
    
    PromotionUrl   string `json:"promotion_url,omitempty" xml:"promotion_url,omitempty"`

    
    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 错误信息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}
