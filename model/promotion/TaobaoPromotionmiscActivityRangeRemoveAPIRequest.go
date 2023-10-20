package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscActivityRangeRemoveAPIRequest 去除活动参与的商品 API请求
// taobao.promotionmisc.activity.range.remove
//
// 去除活动参与的商品
type TaobaoPromotionmiscActivityRangeRemoveAPIRequest struct {
	model.Params
	// 商品id,多个id用逗号隔开。
	_ids string
	// 活动id。
	_activityId int64
}

// NewTaobaoPromotionmiscActivityRangeRemoveRequest 初始化TaobaoPromotionmiscActivityRangeRemoveAPIRequest对象
func NewTaobaoPromotionmiscActivityRangeRemoveRequest() *TaobaoPromotionmiscActivityRangeRemoveAPIRequest {
	return &TaobaoPromotionmiscActivityRangeRemoveAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionmiscActivityRangeRemoveAPIRequest) Reset() {
	r._ids = ""
	r._activityId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscActivityRangeRemoveAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.activity.range.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscActivityRangeRemoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscActivityRangeRemoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIds is Ids Setter
// 商品id,多个id用逗号隔开。
func (r *TaobaoPromotionmiscActivityRangeRemoveAPIRequest) SetIds(_ids string) error {
	r._ids = _ids
	r.Set("ids", _ids)
	return nil
}

// GetIds Ids Getter
func (r TaobaoPromotionmiscActivityRangeRemoveAPIRequest) GetIds() string {
	return r._ids
}

// SetActivityId is ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscActivityRangeRemoveAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionmiscActivityRangeRemoveAPIRequest) GetActivityId() int64 {
	return r._activityId
}

var poolTaobaoPromotionmiscActivityRangeRemoveAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionmiscActivityRangeRemoveRequest()
	},
}

// GetTaobaoPromotionmiscActivityRangeRemoveRequest 从 sync.Pool 获取 TaobaoPromotionmiscActivityRangeRemoveAPIRequest
func GetTaobaoPromotionmiscActivityRangeRemoveAPIRequest() *TaobaoPromotionmiscActivityRangeRemoveAPIRequest {
	return poolTaobaoPromotionmiscActivityRangeRemoveAPIRequest.Get().(*TaobaoPromotionmiscActivityRangeRemoveAPIRequest)
}

// ReleaseTaobaoPromotionmiscActivityRangeRemoveAPIRequest 将 TaobaoPromotionmiscActivityRangeRemoveAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionmiscActivityRangeRemoveAPIRequest(v *TaobaoPromotionmiscActivityRangeRemoveAPIRequest) {
	v.Reset()
	poolTaobaoPromotionmiscActivityRangeRemoveAPIRequest.Put(v)
}
