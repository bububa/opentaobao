package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionmisccommonitemactivitygetAPIRequest 查询通用单品优惠活动 API请求
// taobao.promotionmisc.common.item.activity.get
//
// 查询通用单品优惠活动。
type TaobaopromotionmisccommonitemactivitygetAPIRequest struct {
	model.Params
	// 优惠活动ID
	_activityId int64
}

// NewTaobaopromotionmisccommonitemactivitygetRequest 初始化TaobaopromotionmisccommonitemactivitygetAPIRequest对象
func NewTaobaopromotionmisccommonitemactivitygetRequest() *TaobaopromotionmisccommonitemactivitygetAPIRequest {
	return &TaobaopromotionmisccommonitemactivitygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionmisccommonitemactivitygetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.common.item.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionmisccommonitemactivitygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionmisccommonitemactivitygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 优惠活动ID
func (r *TaobaopromotionmisccommonitemactivitygetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaopromotionmisccommonitemactivitygetAPIRequest) GetActivityId() int64 {
	return r._activityId
}
