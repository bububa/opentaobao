package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询满就送活动 APIResponse
taobao.promotionmisc.mjs.activity.get

查询满就送活动
*/
type TaobaoPromotionmiscMjsActivityGetAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionmiscMjsActivityGetResponse
}

type TaobaoPromotionmiscMjsActivityGetResponse struct {
    XMLName xml.Name `xml:"promotionmisc_mjs_activity_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 满就送活动信息。
    
    MjsPromotion   *MjsPromotion `json:"mjs_promotion,omitempty" xml:"mjs_promotion,omitempty"`

    
}
