package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscItemActivityDeleteAPIRequest 删除无条件单品优惠活动 API请求
// taobao.promotionmisc.item.activity.delete
//
// 删除无条件单品优惠活动
type TaobaoPromotionmiscItemActivityDeleteAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
}

// NewTaobaoPromotionmiscItemActivityDeleteRequest 初始化TaobaoPromotionmiscItemActivityDeleteAPIRequest对象
func NewTaobaoPromotionmiscItemActivityDeleteRequest() *TaobaoPromotionmiscItemActivityDeleteAPIRequest {
	return &TaobaoPromotionmiscItemActivityDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscItemActivityDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.item.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscItemActivityDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscItemActivityDeleteAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// Get ActivityId Getter
func (r TaobaoPromotionmiscItemActivityDeleteAPIRequest) GetActivityId() int64 {
	return r._activityId
}
