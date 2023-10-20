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
	_searchOption *TopApiRefundRptOption
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

// SetSearchOption is SearchOption Setter
// 参数option
func (r *TaobaotbkrelationrefundAPIRequest) SetSearchOption(_searchOption *TopApiRefundRptOption) error {
	r._searchOption = _searchOption
	r.Set("search_option", _searchOption)
	return nil
}

// GetSearchOption SearchOption Getter
func (r TaobaotbkrelationrefundAPIRequest) GetSearchOption() *TopApiRefundRptOption {
	return r._searchOption
}
