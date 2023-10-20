package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscCommonItemActivityListGetAPIRequest 查询通用单品优惠活动列表 API请求
// taobao.promotionmisc.common.item.activity.list.get
//
// 查询通用单品优惠活动列表。
type TaobaoPromotionmiscCommonItemActivityListGetAPIRequest struct {
	model.Params
	// 分页页码，页码从1开始
	_pageNo int64
	// 分页大小，不能超过50
	_pageSize int64
}

// NewTaobaoPromotionmiscCommonItemActivityListGetRequest 初始化TaobaoPromotionmiscCommonItemActivityListGetAPIRequest对象
func NewTaobaoPromotionmiscCommonItemActivityListGetRequest() *TaobaoPromotionmiscCommonItemActivityListGetAPIRequest {
	return &TaobaoPromotionmiscCommonItemActivityListGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionmiscCommonItemActivityListGetAPIRequest) Reset() {
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemActivityListGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.common.item.activity.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemActivityListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscCommonItemActivityListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageNo is PageNo Setter
// 分页页码，页码从1开始
func (r *TaobaoPromotionmiscCommonItemActivityListGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoPromotionmiscCommonItemActivityListGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页大小，不能超过50
func (r *TaobaoPromotionmiscCommonItemActivityListGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoPromotionmiscCommonItemActivityListGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoPromotionmiscCommonItemActivityListGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionmiscCommonItemActivityListGetRequest()
	},
}

// GetTaobaoPromotionmiscCommonItemActivityListGetRequest 从 sync.Pool 获取 TaobaoPromotionmiscCommonItemActivityListGetAPIRequest
func GetTaobaoPromotionmiscCommonItemActivityListGetAPIRequest() *TaobaoPromotionmiscCommonItemActivityListGetAPIRequest {
	return poolTaobaoPromotionmiscCommonItemActivityListGetAPIRequest.Get().(*TaobaoPromotionmiscCommonItemActivityListGetAPIRequest)
}

// ReleaseTaobaoPromotionmiscCommonItemActivityListGetAPIRequest 将 TaobaoPromotionmiscCommonItemActivityListGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionmiscCommonItemActivityListGetAPIRequest(v *TaobaoPromotionmiscCommonItemActivityListGetAPIRequest) {
	v.Reset()
	poolTaobaoPromotionmiscCommonItemActivityListGetAPIRequest.Put(v)
}
