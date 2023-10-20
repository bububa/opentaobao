package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscCommonItemActivityDeleteAPIRequest 删除通用单品优惠活动 API请求
// taobao.promotionmisc.common.item.activity.delete
//
// 删除通用单品优惠活动。
type TaobaoPromotionmiscCommonItemActivityDeleteAPIRequest struct {
	model.Params
	// 优惠活动ID
	_activityId int64
}

// NewTaobaoPromotionmiscCommonItemActivityDeleteRequest 初始化TaobaoPromotionmiscCommonItemActivityDeleteAPIRequest对象
func NewTaobaoPromotionmiscCommonItemActivityDeleteRequest() *TaobaoPromotionmiscCommonItemActivityDeleteAPIRequest {
	return &TaobaoPromotionmiscCommonItemActivityDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemActivityDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.common.item.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemActivityDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscCommonItemActivityDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 优惠活动ID
func (r *TaobaoPromotionmiscCommonItemActivityDeleteAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionmiscCommonItemActivityDeleteAPIRequest) GetActivityId() int64 {
	return r._activityId
}
