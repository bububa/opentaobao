package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取卖家最低折扣 APIResponse
taobao.ump.promotion.global.discount.get

提供卖家最低折扣查询功能
*/
type TaobaoUmpPromotionGlobalDiscountGetAPIResponse struct {
    model.CommonResponse
    TaobaoUmpPromotionGlobalDiscountGetResponse
}

type TaobaoUmpPromotionGlobalDiscountGetResponse struct {
    XMLName xml.Name `xml:"ump_promotion_global_discount_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果对象
    
    Result   *TaobaoUmpPromotionGlobalDiscountGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
