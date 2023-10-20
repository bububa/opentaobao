package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscItemActivityListGetAPIRequest 查询无条件单品优惠活动列表 API请求
// taobao.promotionmisc.item.activity.list.get
//
// 查询无条件单品优惠活动列表
type TaobaoPromotionmiscItemActivityListGetAPIRequest struct {
	model.Params
	// 页码。
	_pageNo int64
	// 每页记录数，最大支持50 。
	_pageSize int64
}

// NewTaobaoPromotionmiscItemActivityListGetRequest 初始化TaobaoPromotionmiscItemActivityListGetAPIRequest对象
func NewTaobaoPromotionmiscItemActivityListGetRequest() *TaobaoPromotionmiscItemActivityListGetAPIRequest {
	return &TaobaoPromotionmiscItemActivityListGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscItemActivityListGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.item.activity.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscItemActivityListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscItemActivityListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageNo is PageNo Setter
// 页码。
func (r *TaobaoPromotionmiscItemActivityListGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoPromotionmiscItemActivityListGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页记录数，最大支持50 。
func (r *TaobaoPromotionmiscItemActivityListGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoPromotionmiscItemActivityListGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
