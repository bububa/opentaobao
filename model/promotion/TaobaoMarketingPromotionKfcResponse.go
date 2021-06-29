package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向优惠活动名称与描述违禁词检查 APIResponse
taobao.marketing.promotion.kfc

活动名称与描述违禁词检查
*/
type TaobaoMarketingPromotionKfcAPIResponse struct {
    model.CommonResponse
    TaobaoMarketingPromotionKfcResponse
}

type TaobaoMarketingPromotionKfcResponse struct {
    XMLName xml.Name `xml:"marketing_promotion_kfc_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
