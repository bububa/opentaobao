package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscCommonItemActivityGetAPIRequest 查询通用单品优惠活动 API请求
// taobao.promotionmisc.common.item.activity.get
//
// 查询通用单品优惠活动。
type TaobaoPromotionmiscCommonItemActivityGetAPIRequest struct {
	model.Params
	// 优惠活动ID
	_activityId int64
}

// NewTaobaoPromotionmiscCommonItemActivityGetRequest 初始化TaobaoPromotionmiscCommonItemActivityGetAPIRequest对象
func NewTaobaoPromotionmiscCommonItemActivityGetRequest() *TaobaoPromotionmiscCommonItemActivityGetAPIRequest {
	return &TaobaoPromotionmiscCommonItemActivityGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemActivityGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.common.item.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemActivityGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetActivityId is ActivityId Setter
// 优惠活动ID
func (r *TaobaoPromotionmiscCommonItemActivityGetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionmiscCommonItemActivityGetAPIRequest) GetActivityId() int64 {
	return r._activityId
}
