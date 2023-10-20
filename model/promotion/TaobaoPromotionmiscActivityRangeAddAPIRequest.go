package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionmiscactivityrangeaddAPIRequest 增加活动参与的商品 API请求
// taobao.promotionmisc.activity.range.add
//
// 增加活动参与的商品，部分商品参与的活动，最大支持指定150个商品。
type TaobaopromotionmiscactivityrangeaddAPIRequest struct {
	model.Params
	// 商品id,多个id用逗号隔开，一次不超过50个。
	_ids string
	// 活动id。
	_activityId int64
}

// NewTaobaopromotionmiscactivityrangeaddRequest 初始化TaobaopromotionmiscactivityrangeaddAPIRequest对象
func NewTaobaopromotionmiscactivityrangeaddRequest() *TaobaopromotionmiscactivityrangeaddAPIRequest {
	return &TaobaopromotionmiscactivityrangeaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionmiscactivityrangeaddAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.activity.range.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionmiscactivityrangeaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionmiscactivityrangeaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIds is Ids Setter
// 商品id,多个id用逗号隔开，一次不超过50个。
func (r *TaobaopromotionmiscactivityrangeaddAPIRequest) SetIds(_ids string) error {
	r._ids = _ids
	r.Set("ids", _ids)
	return nil
}

// GetIds Ids Getter
func (r TaobaopromotionmiscactivityrangeaddAPIRequest) GetIds() string {
	return r._ids
}

// SetActivityId is ActivityId Setter
// 活动id。
func (r *TaobaopromotionmiscactivityrangeaddAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaopromotionmiscactivityrangeaddAPIRequest) GetActivityId() int64 {
	return r._activityId
}
