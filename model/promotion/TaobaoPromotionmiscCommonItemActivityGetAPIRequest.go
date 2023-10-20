package promotion

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionmiscCommonItemActivityGetAPIRequest) Reset() {
	r._activityId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemActivityGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.common.item.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemActivityGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscCommonItemActivityGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoPromotionmiscCommonItemActivityGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionmiscCommonItemActivityGetRequest()
	},
}

// GetTaobaoPromotionmiscCommonItemActivityGetRequest 从 sync.Pool 获取 TaobaoPromotionmiscCommonItemActivityGetAPIRequest
func GetTaobaoPromotionmiscCommonItemActivityGetAPIRequest() *TaobaoPromotionmiscCommonItemActivityGetAPIRequest {
	return poolTaobaoPromotionmiscCommonItemActivityGetAPIRequest.Get().(*TaobaoPromotionmiscCommonItemActivityGetAPIRequest)
}

// ReleaseTaobaoPromotionmiscCommonItemActivityGetAPIRequest 将 TaobaoPromotionmiscCommonItemActivityGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionmiscCommonItemActivityGetAPIRequest(v *TaobaoPromotionmiscCommonItemActivityGetAPIRequest) {
	v.Reset()
	poolTaobaoPromotionmiscCommonItemActivityGetAPIRequest.Put(v)
}
