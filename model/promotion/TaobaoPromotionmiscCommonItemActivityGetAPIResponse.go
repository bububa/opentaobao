package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询通用单品优惠活动 API返回值 
taobao.promotionmisc.common.item.activity.get

查询通用单品优惠活动。
*/
type TaobaoPromotionmiscCommonItemActivityGetAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionmiscCommonItemActivityGetAPIResponseModel
}

// 查询通用单品优惠活动 成功返回结果
type TaobaoPromotionmiscCommonItemActivityGetAPIResponseModel struct {
    XMLName xml.Name `xml:"promotionmisc_common_item_activity_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否查询成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 优惠活动
    Activity   *CommonItemActivity `json:"activity,omitempty" xml:"activity,omitempty"`
}
