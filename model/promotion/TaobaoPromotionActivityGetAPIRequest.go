package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionactivitygetAPIRequest 查询某个卖家的店铺优惠券领取活动 API请求
// taobao.promotion.activity.get
//
// 查询某个卖家的店铺优惠券领取活动&lt;br/&gt;返回，优惠券领取活动ID，优惠券ID，总领用量，每人限领量，已领取数量&lt;br/&gt;领取活动状态，优惠券领取链接&lt;br/&gt;最多50个优惠券
type TaobaopromotionactivitygetAPIRequest struct {
	model.Params
	// 活动的id
	_activityId int64
}

// NewTaobaopromotionactivitygetRequest 初始化TaobaopromotionactivitygetAPIRequest对象
func NewTaobaopromotionactivitygetRequest() *TaobaopromotionactivitygetAPIRequest {
	return &TaobaopromotionactivitygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionactivitygetAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionactivitygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionactivitygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动的id
func (r *TaobaopromotionactivitygetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaopromotionactivitygetAPIRequest) GetActivityId() int64 {
	return r._activityId
}
