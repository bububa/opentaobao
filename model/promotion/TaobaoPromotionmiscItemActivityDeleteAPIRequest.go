package promotion

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionmiscItemActivityDeleteAPIRequest) Reset() {
	r._activityId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscItemActivityDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.item.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscItemActivityDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscItemActivityDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscItemActivityDeleteAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionmiscItemActivityDeleteAPIRequest) GetActivityId() int64 {
	return r._activityId
}

var poolTaobaoPromotionmiscItemActivityDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionmiscItemActivityDeleteRequest()
	},
}

// GetTaobaoPromotionmiscItemActivityDeleteRequest 从 sync.Pool 获取 TaobaoPromotionmiscItemActivityDeleteAPIRequest
func GetTaobaoPromotionmiscItemActivityDeleteAPIRequest() *TaobaoPromotionmiscItemActivityDeleteAPIRequest {
	return poolTaobaoPromotionmiscItemActivityDeleteAPIRequest.Get().(*TaobaoPromotionmiscItemActivityDeleteAPIRequest)
}

// ReleaseTaobaoPromotionmiscItemActivityDeleteAPIRequest 将 TaobaoPromotionmiscItemActivityDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionmiscItemActivityDeleteAPIRequest(v *TaobaoPromotionmiscItemActivityDeleteAPIRequest) {
	v.Reset()
	poolTaobaoPromotionmiscItemActivityDeleteAPIRequest.Put(v)
}
