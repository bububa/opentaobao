package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询某个卖家的店铺优惠券领取活动 API返回值 
taobao.promotion.activity.get

查询某个卖家的店铺优惠券领取活动<br/>返回，优惠券领取活动ID，优惠券ID，总领用量，每人限领量，已领取数量<br/>领取活动状态，优惠券领取链接<br/>最多50个优惠券
*/
type TaobaoPromotionActivityGetAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionActivityGetResponse
}

// 查询某个卖家的店铺优惠券领取活动 成功返回结果
type TaobaoPromotionActivityGetResponse struct {
    XMLName xml.Name `xml:"promotion_activity_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 活动列表
    Activitys   []Activity `json:"activitys,omitempty" xml:"activitys>activity,omitempty"`
}
