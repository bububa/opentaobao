package promotion

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest) Reset() {
	r._activityId = 0
	r._detailId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.common.item.detail.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 优惠活动ID
func (r *TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetDetailId is DetailId Setter
// 优惠详情ID
func (r *TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest) SetDetailId(_detailId int64) error {
	r._detailId = _detailId
	r.Set("detail_id", _detailId)
	return nil
}

// GetDetailId DetailId Getter
func (r TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest) GetDetailId() int64 {
	return r._detailId
}

var poolTaobaoPromotionmiscCommonItemDetailDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionmiscCommonItemDetailDeleteRequest()
	},
}

// GetTaobaoPromotionmiscCommonItemDetailDeleteRequest 从 sync.Pool 获取 TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest
func GetTaobaoPromotionmiscCommonItemDetailDeleteAPIRequest() *TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest {
	return poolTaobaoPromotionmiscCommonItemDetailDeleteAPIRequest.Get().(*TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest)
}

// ReleaseTaobaoPromotionmiscCommonItemDetailDeleteAPIRequest 将 TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionmiscCommonItemDetailDeleteAPIRequest(v *TaobaoPromotionmiscCommonItemDetailDeleteAPIRequest) {
	v.Reset()
	poolTaobaoPromotionmiscCommonItemDetailDeleteAPIRequest.Put(v)
}
