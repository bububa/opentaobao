package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscCommonItemDetailListGetAPIRequest 查询通用单品优惠详情列表 API请求
// taobao.promotionmisc.common.item.detail.list.get
//
// 查询通用单品优惠详情列表。
type TaobaoPromotionmiscCommonItemDetailListGetAPIRequest struct {
	model.Params
	// 优惠活动ID
	_activityId int64
	// 分页页码，页码从1开始
	_pageNo int64
	// 分页大小，不能超过50
	_pageSize int64
}

// NewTaobaoPromotionmiscCommonItemDetailListGetRequest 初始化TaobaoPromotionmiscCommonItemDetailListGetAPIRequest对象
func NewTaobaoPromotionmiscCommonItemDetailListGetRequest() *TaobaoPromotionmiscCommonItemDetailListGetAPIRequest {
	return &TaobaoPromotionmiscCommonItemDetailListGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionmiscCommonItemDetailListGetAPIRequest) Reset() {
	r._activityId = 0
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemDetailListGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.common.item.detail.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemDetailListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscCommonItemDetailListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 优惠活动ID
func (r *TaobaoPromotionmiscCommonItemDetailListGetAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionmiscCommonItemDetailListGetAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetPageNo is PageNo Setter
// 分页页码，页码从1开始
func (r *TaobaoPromotionmiscCommonItemDetailListGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoPromotionmiscCommonItemDetailListGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页大小，不能超过50
func (r *TaobaoPromotionmiscCommonItemDetailListGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoPromotionmiscCommonItemDetailListGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoPromotionmiscCommonItemDetailListGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionmiscCommonItemDetailListGetRequest()
	},
}

// GetTaobaoPromotionmiscCommonItemDetailListGetRequest 从 sync.Pool 获取 TaobaoPromotionmiscCommonItemDetailListGetAPIRequest
func GetTaobaoPromotionmiscCommonItemDetailListGetAPIRequest() *TaobaoPromotionmiscCommonItemDetailListGetAPIRequest {
	return poolTaobaoPromotionmiscCommonItemDetailListGetAPIRequest.Get().(*TaobaoPromotionmiscCommonItemDetailListGetAPIRequest)
}

// ReleaseTaobaoPromotionmiscCommonItemDetailListGetAPIRequest 将 TaobaoPromotionmiscCommonItemDetailListGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionmiscCommonItemDetailListGetAPIRequest(v *TaobaoPromotionmiscCommonItemDetailListGetAPIRequest) {
	v.Reset()
	poolTaobaoPromotionmiscCommonItemDetailListGetAPIRequest.Put(v)
}
