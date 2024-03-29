package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest 清空活动参与的商品 API请求
// taobao.promotionmisc.activity.range.all.remove
//
// 清空活动参与的商品
type TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
}

// NewTaobaoPromotionmiscActivityRangeAllRemoveRequest 初始化TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest对象
func NewTaobaoPromotionmiscActivityRangeAllRemoveRequest() *TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest {
	return &TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest) Reset() {
	r._activityId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.activity.range.all.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest) GetActivityId() int64 {
	return r._activityId
}

var poolTaobaoPromotionmiscActivityRangeAllRemoveAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionmiscActivityRangeAllRemoveRequest()
	},
}

// GetTaobaoPromotionmiscActivityRangeAllRemoveRequest 从 sync.Pool 获取 TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest
func GetTaobaoPromotionmiscActivityRangeAllRemoveAPIRequest() *TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest {
	return poolTaobaoPromotionmiscActivityRangeAllRemoveAPIRequest.Get().(*TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest)
}

// ReleaseTaobaoPromotionmiscActivityRangeAllRemoveAPIRequest 将 TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionmiscActivityRangeAllRemoveAPIRequest(v *TaobaoPromotionmiscActivityRangeAllRemoveAPIRequest) {
	v.Reset()
	poolTaobaoPromotionmiscActivityRangeAllRemoveAPIRequest.Put(v)
}
