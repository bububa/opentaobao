package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionActivityGetAPIRequest 查询某个卖家的店铺优惠券领取活动 API请求
// taobao.promotion.activity.get
//
// 查询某个卖家的店铺优惠券领取活动<br/>返回，优惠券领取活动ID，优惠券ID，总领用量，每人限领量，已领取数量<br/>领取活动状态，优惠券领取链接<br/>最多50个优惠券
type TaobaoPromotionActivityGetAPIRequest struct {
	model.Params
	// 活动的id
	_activityId int64
}

// NewTaobaoPromotionActivityGetRequest 初始化TaobaoPromotionActivityGetAPIRequest对象
func NewTaobaoPromotionActivityGetRequest() *TaobaoPromotionActivityGetAPIRequest {
	return &TaobaoPromotionActivityGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionActivityGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionActivityGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetActivityId is ActivityId Setter
// 活动的id
func (r *TaobaoPromotionActivityGetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionActivityGetAPIRequest) GetActivityId() int64 {
	return r._activityId
}
