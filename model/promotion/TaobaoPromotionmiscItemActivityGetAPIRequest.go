package promotion

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionmiscItemActivityGetAPIRequest) Reset() {
	r._activityId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscItemActivityGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.item.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscItemActivityGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscItemActivityGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscItemActivityGetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionmiscItemActivityGetAPIRequest) GetActivityId() int64 {
	return r._activityId
}

var poolTaobaoPromotionmiscItemActivityGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionmiscItemActivityGetRequest()
	},
}

// GetTaobaoPromotionmiscItemActivityGetRequest 从 sync.Pool 获取 TaobaoPromotionmiscItemActivityGetAPIRequest
func GetTaobaoPromotionmiscItemActivityGetAPIRequest() *TaobaoPromotionmiscItemActivityGetAPIRequest {
	return poolTaobaoPromotionmiscItemActivityGetAPIRequest.Get().(*TaobaoPromotionmiscItemActivityGetAPIRequest)
}

// ReleaseTaobaoPromotionmiscItemActivityGetAPIRequest 将 TaobaoPromotionmiscItemActivityGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionmiscItemActivityGetAPIRequest(v *TaobaoPromotionmiscItemActivityGetAPIRequest) {
	v.Reset()
	poolTaobaoPromotionmiscItemActivityGetAPIRequest.Put(v)
}
