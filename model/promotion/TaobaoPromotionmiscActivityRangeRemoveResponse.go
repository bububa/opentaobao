package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
去除活动参与的商品 API返回值 
taobao.promotionmisc.activity.range.remove

去除活动参与的商品
*/
type TaobaoPromotionmiscActivityRangeRemoveAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionmiscActivityRangeRemoveResponse
}

// 去除活动参与的商品 成功返回结果
type TaobaoPromotionmiscActivityRangeRemoveResponse struct {
    XMLName xml.Name `xml:"promotionmisc_activity_range_remove_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 去除活动参与的商品是否成功。
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
