package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionmiscmjsactivitygetAPIRequest 查询满就送活动 API请求
// taobao.promotionmisc.mjs.activity.get
//
// 查询满就送活动
type TaobaopromotionmiscmjsactivitygetAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
}

// NewTaobaopromotionmiscmjsactivitygetRequest 初始化TaobaopromotionmiscmjsactivitygetAPIRequest对象
func NewTaobaopromotionmiscmjsactivitygetRequest() *TaobaopromotionmiscmjsactivitygetAPIRequest {
	return &TaobaopromotionmiscmjsactivitygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionmiscmjsactivitygetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.mjs.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionmiscmjsactivitygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionmiscmjsactivitygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id。
func (r *TaobaopromotionmiscmjsactivitygetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaopromotionmiscmjsactivitygetAPIRequest) GetActivityId() int64 {
	return r._activityId
}
