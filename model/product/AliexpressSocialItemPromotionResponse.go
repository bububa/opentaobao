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
	RequestId     string         `json:"request_id,omitempty" xml:"aliexpress_social_item_promotion_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 推广链接
    
    PromotionUrl   string `json:"promotion_url,omitempty" xml:"