package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionmiscactivityrangeallremoveAPIRequest 清空活动参与的商品 API请求
// taobao.promotionmisc.activity.range.all.remove
//
// 清空活动参与的商品
type TaobaopromotionmiscactivityrangeallremoveAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
}

// NewTaobaopromotionmiscactivityrangeallremoveRequest 初始化TaobaopromotionmiscactivityrangeallremoveAPIRequest对象
func NewTaobaopromotionmiscactivityrangeallremoveRequest() *TaobaopromotionmiscactivityrangeallremoveAPIRequest {
	return &TaobaopromotionmiscactivityrangeallremoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionmiscactivityrangeallremoveAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.activity.range.all.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionmiscactivityrangeallremoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionmiscactivityrangeallremoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id。
func (r *TaobaopromotionmiscactivityrangeallremoveAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaopromotionmiscactivityrangeallremoveAPIRequest) GetActivityId() int64 {
	return r._activityId
}
