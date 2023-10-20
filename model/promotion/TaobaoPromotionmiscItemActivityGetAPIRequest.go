package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionmiscitemactivitygetAPIRequest 查询无条件单品优惠活动 API请求
// taobao.promotionmisc.item.activity.get
//
// 查询无条件单品优惠活动
type TaobaopromotionmiscitemactivitygetAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
}

// NewTaobaopromotionmiscitemactivitygetRequest 初始化TaobaopromotionmiscitemactivitygetAPIRequest对象
func NewTaobaopromotionmiscitemactivitygetRequest() *TaobaopromotionmiscitemactivitygetAPIRequest {
	return &TaobaopromotionmiscitemactivitygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionmiscitemactivitygetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.item.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionmiscitemactivitygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionmiscitemactivitygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id。
func (r *TaobaopromotionmiscitemactivitygetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaopromotionmiscitemactivitygetAPIRequest) GetActivityId() int64 {
	return r._activityId
}
