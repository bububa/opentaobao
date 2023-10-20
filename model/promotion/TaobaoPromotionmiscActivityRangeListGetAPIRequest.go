package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscActivityRangeListGetAPIRequest 查询活动参与的商品 API请求
// taobao.promotionmisc.activity.range.list.get
//
// 查询活动参与的商品
type TaobaoPromotionmiscActivityRangeListGetAPIRequest struct {
	model.Params
	// 活动id
	_activityId int64
}

// NewTaobaoPromotionmiscActivityRangeListGetRequest 初始化TaobaoPromotionmiscActivityRangeListGetAPIRequest对象
func NewTaobaoPromotionmiscActivityRangeListGetRequest() *TaobaoPromotionmiscActivityRangeListGetAPIRequest {
	return &TaobaoPromotionmiscActivityRangeListGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscActivityRangeListGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.activity.range.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscActivityRangeListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscActivityRangeListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *TaobaoPromotionmiscActivityRangeListGetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionmiscActivityRangeListGetAPIRequest) GetActivityId() int64 {
	return r._activityId
}
