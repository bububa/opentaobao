package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscMjsActivityGetAPIRequest 查询满就送活动 API请求
// taobao.promotionmisc.mjs.activity.get
//
// 查询满就送活动
type TaobaoPromotionmiscMjsActivityGetAPIRequest struct {
	model.Params
	// 活动id。
	_activityId int64
}

// NewTaobaoPromotionmiscMjsActivityGetRequest 初始化TaobaoPromotionmiscMjsActivityGetAPIRequest对象
func NewTaobaoPromotionmiscMjsActivityGetRequest() *TaobaoPromotionmiscMjsActivityGetAPIRequest {
	return &TaobaoPromotionmiscMjsActivityGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionmiscMjsActivityGetAPIRequest) Reset() {
	r._activityId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscMjsActivityGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.mjs.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscMjsActivityGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscMjsActivityGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscMjsActivityGetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionmiscMjsActivityGetAPIRequest) GetActivityId() int64 {
	return r._activityId
}

var poolTaobaoPromotionmiscMjsActivityGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionmiscMjsActivityGetRequest()
	},
}

// GetTaobaoPromotionmiscMjsActivityGetRequest 从 sync.Pool 获取 TaobaoPromotionmiscMjsActivityGetAPIRequest
func GetTaobaoPromotionmiscMjsActivityGetAPIRequest() *TaobaoPromotionmiscMjsActivityGetAPIRequest {
	return poolTaobaoPromotionmiscMjsActivityGetAPIRequest.Get().(*TaobaoPromotionmiscMjsActivityGetAPIRequest)
}

// ReleaseTaobaoPromotionmiscMjsActivityGetAPIRequest 将 TaobaoPromotionmiscMjsActivityGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionmiscMjsActivityGetAPIRequest(v *TaobaoPromotionmiscMjsActivityGetAPIRequest) {
	v.Reset()
	poolTaobaoPromotionmiscMjsActivityGetAPIRequest.Put(v)
}
