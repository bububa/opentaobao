package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderitemPageGetAPIRequest 分页查询物流宝订单商品详情 API请求
// taobao.wlb.orderitem.page.get
//
// 分页查询物流宝订单商品详情
type TaobaoWlbOrderitemPageGetAPIRequest struct {
	model.Params
	// 物流宝订单编码
	_orderCode string
	// 分页查询参数，指定查询页数，默认为1
	_pageNo int64
	// 分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询
	_pageSize int64
}

// NewTaobaoWlbOrderitemPageGetRequest 初始化TaobaoWlbOrderitemPageGetAPIRequest对象
func NewTaobaoWlbOrderitemPageGetRequest() *TaobaoWlbOrderitemPageGetAPIRequest {
	return &TaobaoWlbOrderitemPageGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderitemPageGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.orderitem.page.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderitemPageGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderCode is OrderCode Setter
// 物流宝订单编码
func (r *TaobaoWlbOrderitemPageGetAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaoWlbOrderitemPageGetAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetPageNo is PageNo Setter
// 分页查询参数，指定查询页数，默认为1
func (r *TaobaoWlbOrderitemPageGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoWlbOrderitemPageGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询
func (r *TaobaoWlbOrderitemPageGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoWlbOrderitemPageGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
