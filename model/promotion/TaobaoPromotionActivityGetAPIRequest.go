package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionActivityGetAPIRequest
查询某个卖家的店铺优惠券领取活动 API请求
taobao.promotion.activity.get

查询某个卖家的店铺优惠券领取活动<br/>返回，优惠券领取活动ID，优惠券ID，总领用量，每人限领量，已领取数量<br/>领取活动状态，优惠券领取链接<br/>最多50个优惠券 */
type TaobaoPromotionActivityGetAPIRequest struct {
	model.Params
	// 活动的id
	_activityId int64
}

// New
