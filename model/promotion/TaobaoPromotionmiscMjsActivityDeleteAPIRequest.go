package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscMjsActivityDeleteAPIRequest 删除满就送活动 API请求
// taobao.promotionmisc.mjs.activity.delete
//
// 删除满就送活动
type TaobaoPromotionmiscMjsActivityDeleteAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
}

// NewTaobaoPromotionmiscMjsActivityDeleteRequest 初始化TaobaoPromotionmiscMjsActivityDeleteAPIRequest对象
func NewTaobaoPromotionmiscMjsActivityDeleteRequest() *TaobaoPromotionmiscMjsActivityDeleteAPIRequest {
	return &TaobaoPromotionmiscMjsActivityDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionmiscMjsActivityDeleteAPIRequest) Reset() {
	r._activityId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscMjsActivityDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.mjs.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscMjsActivityDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscMjsActivityDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscMjsActivityDeleteAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionmiscMjsActivityDeleteAPIRequest) GetActivityId() int64 {
	return r._activityId
}

var poolTaobaoPromotionmiscMjsActivityDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionmiscMjsActivityDeleteRequest()
	},
}

// GetTaobaoPromotionmiscMjsActivityDeleteRequest 从 sync.Pool 获取 TaobaoPromotionmiscMjsActivityDeleteAPIRequest
func GetTaobaoPromotionmiscMjsActivityDeleteAPIRequest() *TaobaoPromotionmiscMjsActivityDeleteAPIRequest {
	return poolTaobaoPromotionmiscMjsActivityDeleteAPIRequest.Get().(*TaobaoPromotionmiscMjsActivityDeleteAPIRequest)
}

// ReleaseTaobaoPromotionmiscMjsActivityDeleteAPIRequest 将 TaobaoPromotionmiscMjsActivityDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionmiscMjsActivityDeleteAPIRequest(v *TaobaoPromotionmiscMjsActivityDeleteAPIRequest) {
	v.Reset()
	poolTaobaoPromotionmiscMjsActivityDeleteAPIRequest.Put(v)
}
