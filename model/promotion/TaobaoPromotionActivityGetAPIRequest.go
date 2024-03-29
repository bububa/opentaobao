package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionActivityGetAPIRequest 查询某个卖家的店铺优惠券领取活动 API请求
// taobao.promotion.activity.get
//
// 查询某个卖家的店铺优惠券领取活动&lt;br/&gt;返回，优惠券领取活动ID，优惠券ID，总领用量，每人限领量，已领取数量&lt;br/&gt;领取活动状态，优惠券领取链接&lt;br/&gt;最多50个优惠券
type TaobaoPromotionActivityGetAPIRequest struct {
	model.Params
	// 活动的id
	_activityId int64
}

// NewTaobaoPromotionActivityGetRequest 初始化TaobaoPromotionActivityGetAPIRequest对象
func NewTaobaoPromotionActivityGetRequest() *TaobaoPromotionActivityGetAPIRequest {
	return &TaobaoPromotionActivityGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionActivityGetAPIRequest) Reset() {
	r._activityId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionActivityGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionActivityGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionActivityGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动的id
func (r *TaobaoPromotionActivityGetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionActivityGetAPIRequest) GetActivityId() int64 {
	return r._activityId
}

var poolTaobaoPromotionActivityGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionActivityGetRequest()
	},
}

// GetTaobaoPromotionActivityGetRequest 从 sync.Pool 获取 TaobaoPromotionActivityGetAPIRequest
func GetTaobaoPromotionActivityGetAPIRequest() *TaobaoPromotionActivityGetAPIRequest {
	return poolTaobaoPromotionActivityGetAPIRequest.Get().(*TaobaoPromotionActivityGetAPIRequest)
}

// ReleaseTaobaoPromotionActivityGetAPIRequest 将 TaobaoPromotionActivityGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionActivityGetAPIRequest(v *TaobaoPromotionActivityGetAPIRequest) {
	v.Reset()
	poolTaobaoPromotionActivityGetAPIRequest.Put(v)
}
