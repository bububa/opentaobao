package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscItemActivityGetAPIRequest 查询无条件单品优惠活动 API请求
// taobao.promotionmisc.item.activity.get
//
// 查询无条件单品优惠活动
type TaobaoPromotionmiscItemActivityGetAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
}

// NewTaobaoPromotionmiscItemActivityGetRequest 初始化TaobaoPromotionmiscItemActivityGetAPIRequest对象
func NewTaobaoPromotionmiscItemActivityGetRequest() *TaobaoPromotionmiscItemActivityGetAPIRequest {
	return &TaobaoPromotionmiscItemActivityGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscItemActivityGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.item.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscItemActivityGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscItemActivityGetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// Get ActivityId Getter
func (r TaobaoPromotionmiscItemActivityGetAPIRequest) GetActivityId() int64 {
	return r._activityId
}
