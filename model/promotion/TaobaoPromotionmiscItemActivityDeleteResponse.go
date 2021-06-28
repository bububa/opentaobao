package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除无条件单品优惠活动 APIResponse
taobao.promotionmisc.item.activity.delete

删除无条件单品优惠活动
*/
type TaobaoPromotionmiscItemActivityDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionmiscItemActivityDeleteResponse
}

type TaobaoPromotionmiscItemActivityDeleteResponse struct {
    XMLName xml.Name `xml:"promotionmisc_item_activity_delete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功删除活动。
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
