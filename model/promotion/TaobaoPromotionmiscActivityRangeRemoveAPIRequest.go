package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionmiscactivityrangeremoveAPIRequest 去除活动参与的商品 API请求
// taobao.promotionmisc.activity.range.remove
//
// 去除活动参与的商品
type TaobaopromotionmiscactivityrangeremoveAPIRequest struct {
	model.Params
	// 商品id,多个id用逗号隔开。
	_ids string
	// 活动id。
	_activityId int64
}

// NewTaobaopromotionmiscactivityrangeremoveRequest 初始化TaobaopromotionmiscactivityrangeremoveAPIRequest对象
func NewTaobaopromotionmiscactivityrangeremoveRequest() *TaobaopromotionmiscactivityrangeremoveAPIRequest {
	return &TaobaopromotionmiscactivityrangeremoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionmiscactivityrangeremoveAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.activity.range.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionmiscactivityrangeremoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionmiscactivityrangeremoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIds is Ids Setter
// 商品id,多个id用逗号隔开。
func (r *TaobaopromotionmiscactivityrangeremoveAPIRequest) SetIds(_ids string) error {
	r._ids = _ids
	r.Set("ids", _ids)
	return nil
}

// GetIds Ids Getter
func (r TaobaopromotionmiscactivityrangeremoveAPIRequest) GetIds() string {
	return r._ids
}

// SetActivityId is ActivityId Setter
// 活动id。
func (r *TaobaopromotionmiscactivityrangeremoveAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaopromotionmiscactivityrangeremoveAPIRequest) GetActivityId() int64 {
	return r._activityId
}
