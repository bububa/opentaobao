package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscActivityRangeAddAPIRequest 增加活动参与的商品 API请求
// taobao.promotionmisc.activity.range.add
//
// 增加活动参与的商品，部分商品参与的活动，最大支持指定150个商品。
type TaobaoPromotionmiscActivityRangeAddAPIRequest struct {
	model.Params
	// 商品id,多个id用逗号隔开，一次不超过50个。
	_ids string
	// 活动id。
	_activityId int64
}

// NewTaobaoPromotionmiscActivityRangeAddRequest 初始化TaobaoPromotionmiscActivityRangeAddAPIRequest对象
func NewTaobaoPromotionmiscActivityRangeAddRequest() *TaobaoPromotionmiscActivityRangeAddAPIRequest {
	return &TaobaoPromotionmiscActivityRangeAddAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionmiscActivityRangeAddAPIRequest) Reset() {
	r._ids = ""
	r._activityId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscActivityRangeAddAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.activity.range.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscActivityRangeAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscActivityRangeAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIds is Ids Setter
// 商品id,多个id用逗号隔开，一次不超过50个。
func (r *TaobaoPromotionmiscActivityRangeAddAPIRequest) SetIds(_ids string) error {
	r._ids = _ids
	r.Set("ids", _ids)
	return nil
}

// GetIds Ids Getter
func (r TaobaoPromotionmiscActivityRangeAddAPIRequest) GetIds() string {
	return r._ids
}

// SetActivityId is ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscActivityRangeAddAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionmiscActivityRangeAddAPIRequest) GetActivityId() int64 {
	return r._activityId
}

var poolTaobaoPromotionmiscActivityRangeAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionmiscActivityRangeAddRequest()
	},
}

// GetTaobaoPromotionmiscActivityRangeAddRequest 从 sync.Pool 获取 TaobaoPromotionmiscActivityRangeAddAPIRequest
func GetTaobaoPromotionmiscActivityRangeAddAPIRequest() *TaobaoPromotionmiscActivityRangeAddAPIRequest {
	return poolTaobaoPromotionmiscActivityRangeAddAPIRequest.Get().(*TaobaoPromotionmiscActivityRangeAddAPIRequest)
}

// ReleaseTaobaoPromotionmiscActivityRangeAddAPIRequest 将 TaobaoPromotionmiscActivityRangeAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionmiscActivityRangeAddAPIRequest(v *TaobaoPromotionmiscActivityRangeAddAPIRequest) {
	v.Reset()
	poolTaobaoPromotionmiscActivityRangeAddAPIRequest.Put(v)
}
