package promotion

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscCommonItemActivityListGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.common.item.activity.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscCommonItemActivityListGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
