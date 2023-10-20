package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionmisccommonitemactivitydeleteAPIRequest 删除通用单品优惠活动 API请求
// taobao.promotionmisc.common.item.activity.delete
//
// 删除通用单品优惠活动。
type TaobaopromotionmisccommonitemactivitydeleteAPIRequest struct {
	model.Params
	// 优惠活动ID
	_activityId int64
}

// NewTaobaopromotionmisccommonitemactivitydeleteRequest 初始化TaobaopromotionmisccommonitemactivitydeleteAPIRequest对象
func NewTaobaopromotionmisccommonitemactivitydeleteRequest() *TaobaopromotionmisccommonitemactivitydeleteAPIRequest {
	return &TaobaopromotionmisccommonitemactivitydeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionmisccommonitemactivitydeleteAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.common.item.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionmisccommonitemactivitydeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionmisccommonitemactivitydeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 优惠活动ID
func (r *TaobaopromotionmisccommonitemactivitydeleteAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaopromotionmisccommonitemactivitydeleteAPIRequest) GetActivityId() int64 {
	return r._activityId
}
