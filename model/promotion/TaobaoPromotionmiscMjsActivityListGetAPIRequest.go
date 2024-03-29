package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscMjsActivityListGetAPIRequest 查询满就送活动列表 API请求
// taobao.promotionmisc.mjs.activity.list.get
//
// 查询满就送活动列表。注意，该接口的返回值中，只包含活动的主要信息，如activity_id，name，description，start_time，end_time，type，participate_range。优惠的详细信息，请通过taobao.promotionmisc.mjs.activity.get获取。
type TaobaoPromotionmiscMjsActivityListGetAPIRequest struct {
	model.Params
	// 活动类型： 1表示商品级别的活动；2表示店铺级别的活动。
	_activityType int64
	// 页码。
	_pageNo int64
	// 每页记录数，最大支持50 。
	_pageSize int64
}

// NewTaobaoPromotionmiscMjsActivityListGetRequest 初始化TaobaoPromotionmiscMjsActivityListGetAPIRequest对象
func NewTaobaoPromotionmiscMjsActivityListGetRequest() *TaobaoPromotionmiscMjsActivityListGetAPIRequest {
	return &TaobaoPromotionmiscMjsActivityListGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionmiscMjsActivityListGetAPIRequest) Reset() {
	r._activityType = 0
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscMjsActivityListGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.mjs.activity.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscMjsActivityListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscMjsActivityListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityType is ActivityType Setter
// 活动类型： 1表示商品级别的活动；2表示店铺级别的活动。
func (r *TaobaoPromotionmiscMjsActivityListGetAPIRequest) SetActivityType(_activityType int64) error {
	r._activityType = _activityType
	r.Set("activity_type", _activityType)
	return nil
}

// GetActivityType ActivityType Getter
func (r TaobaoPromotionmiscMjsActivityListGetAPIRequest) GetActivityType() int64 {
	return r._activityType
}

// SetPageNo is PageNo Setter
// 页码。
func (r *TaobaoPromotionmiscMjsActivityListGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoPromotionmiscMjsActivityListGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页记录数，最大支持50 。
func (r *TaobaoPromotionmiscMjsActivityListGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoPromotionmiscMjsActivityListGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoPromotionmiscMjsActivityListGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionmiscMjsActivityListGetRequest()
	},
}

// GetTaobaoPromotionmiscMjsActivityListGetRequest 从 sync.Pool 获取 TaobaoPromotionmiscMjsActivityListGetAPIRequest
func GetTaobaoPromotionmiscMjsActivityListGetAPIRequest() *TaobaoPromotionmiscMjsActivityListGetAPIRequest {
	return poolTaobaoPromotionmiscMjsActivityListGetAPIRequest.Get().(*TaobaoPromotionmiscMjsActivityListGetAPIRequest)
}

// ReleaseTaobaoPromotionmiscMjsActivityListGetAPIRequest 将 TaobaoPromotionmiscMjsActivityListGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionmiscMjsActivityListGetAPIRequest(v *TaobaoPromotionmiscMjsActivityListGetAPIRequest) {
	v.Reset()
	poolTaobaoPromotionmiscMjsActivityListGetAPIRequest.Put(v)
}
