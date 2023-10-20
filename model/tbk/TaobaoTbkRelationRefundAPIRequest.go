package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkrelationrefundAPIRequest 淘宝客-推广者-维权退款订单查询 API请求
// taobao.tbk.relation.refund
//
// 淘宝客维权退款订单查询-渠道管理和会员运营管理专用
type TaobaotbkrelationrefundAPIRequest struct {
	model.Params
	// 参数option
	_searchoption *TopApiRefundRptOption
}

// NewTaobaotbkrelationrefundRequest 初始化TaobaotbkrelationrefundAPIRequest对象
func NewTaobaotbkrelationrefundRequest() *TaobaotbkrelationrefundAPIRequest {
	return &TaobaotbkrelationrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkrelationrefundAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.relation.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkrelationrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkrelationrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSearchoption is Searchoption Setter
// 参数option
func (r *TaobaotbkrelationrefundAPIRequest) SetSearchoption(_searchoption *TopApiRefundRptOption) error {
	r._searchoption = _searchoption
	r.Set("search_option", _searchoption)
	return nil
}

// GetSearchoption Searchoption Getter
func (r TaobaotbkrelationrefundAPIRequest) GetSearchoption() *TopApiRefundRptOption {
	return r._searchoption
}
