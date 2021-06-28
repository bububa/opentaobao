package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询某个卖家的店铺优惠券领取活动 APIResponse
taobao.promotion.activity.get

查询某个卖家的店铺优惠券领取活动<br/>返回，优惠券领取活动ID，优惠券ID，总领用量，每人限领量，已领取数量<br/>领取活动状态，优惠券领取链接<br/>最多50个优惠券
*/
type TaobaoPromotionActivityGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"promotion_activity_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 活动列表
    
    Activitys   []Activity `json:"activitys,omitempty" xml:"