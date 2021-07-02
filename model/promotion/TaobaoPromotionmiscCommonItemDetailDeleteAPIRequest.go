package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest 删除通用单品优惠详情 API请求
// taobao.promotionmisc.common.item.detail.delete
//
// 删除通用单品优惠详情。
type TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest struct {
	model.Params
	// 优惠活动ID
	_activityId int64
	// 优惠详情ID
	_detailId int64
}

// NewTaobaoPromotionmiscCommonItemDetailDeleteRequest 初始化TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest对象
func NewTaobaoPromotionmiscCommonItemDetailDeleteRequest() *TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest {
	return &TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.common.item.detail.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActivityId Setter
// 优惠活动ID
func (r *TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// Get ActivityId Getter
func (r TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// Set is DetailId Setter
// 优惠详情ID
func (r *TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest) SetDetailId(_detailId int64) error {
	r._detailId = _detailId
	r.Set("detail_id", _detailId)
	return nil
}

// Get DetailId Getter
func (r TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest) GetDetailId() int64 {
	return r._detailId
}
