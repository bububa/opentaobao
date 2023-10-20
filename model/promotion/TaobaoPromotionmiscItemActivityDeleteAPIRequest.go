package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionmiscitemactivitydeleteAPIRequest 删除无条件单品优惠活动 API请求
// taobao.promotionmisc.item.activity.delete
//
// 删除无条件单品优惠活动
type TaobaopromotionmiscitemactivitydeleteAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
}

// NewTaobaopromotionmiscitemactivitydeleteRequest 初始化TaobaopromotionmiscitemactivitydeleteAPIRequest对象
func NewTaobaopromotionmiscitemactivitydeleteRequest() *TaobaopromotionmiscitemactivitydeleteAPIRequest {
	return &TaobaopromotionmiscitemactivitydeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionmiscitemactivitydeleteAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.item.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionmiscitemactivitydeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionmiscitemactivitydeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id。
func (r *TaobaopromotionmiscitemactivitydeleteAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaopromotionmiscitemactivitydeleteAPIRequest) GetActivityId() int64 {
	return r._activityId
}
