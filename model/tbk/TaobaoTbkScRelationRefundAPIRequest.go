package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscrelationrefundAPIRequest 淘宝客-服务商-维权退款订单查询 API请求
// taobao.tbk.sc.relation.refund
//
// 淘宝客维权退款订单查询-渠道管理和会员运营管理专用
type TaobaotbkscrelationrefundAPIRequest struct {
	model.Params
	// 参数option
	_searchoption *TopApiRefundRptOption
}

// NewTaobaotbkscrelationrefundRequest 初始化TaobaotbkscrelationrefundAPIRequest对象
func NewTaobaotbkscrelationrefundRequest() *TaobaotbkscrelationrefundAPIRequest {
	return &TaobaotbkscrelationrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscrelationrefundAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.relation.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscrelationrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscrelationrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSearchoption is Searchoption Setter
// 参数option
func (r *TaobaotbkscrelationrefundAPIRequest) SetSearchoption(_searchoption *TopApiRefundRptOption) error {
	r._searchoption = _searchoption
	r.Set("search_option", _searchoption)
	return nil
}

// GetSearchoption Searchoption Getter
func (r TaobaotbkscrelationrefundAPIRequest) GetSearchoption() *TopApiRefundRptOption {
	return r._searchoption
}
