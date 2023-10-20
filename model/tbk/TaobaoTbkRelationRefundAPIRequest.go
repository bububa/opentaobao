package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkRelationRefundAPIRequest 淘宝客-推广者-维权退款订单查询 API请求
// taobao.tbk.relation.refund
//
// 淘宝客维权退款订单查询-渠道管理和会员运营管理专用
type TaobaoTbkRelationRefundAPIRequest struct {
	model.Params
	// 参数option
	_searchOption *TopApiRefundRptOption
}

// NewTaobaoTbkRelationRefundRequest 初始化TaobaoTbkRelationRefundAPIRequest对象
func NewTaobaoTbkRelationRefundRequest() *TaobaoTbkRelationRefundAPIRequest {
	return &TaobaoTbkRelationRefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkRelationRefundAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.relation.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkRelationRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkRelationRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSearchOption is SearchOption Setter
// 参数option
func (r *TaobaoTbkRelationRefundAPIRequest) SetSearchOption(_searchOption *TopApiRefundRptOption) error {
	r._searchOption = _searchOption
	r.Set("search_option", _searchOption)
	return nil
}

// GetSearchOption SearchOption Getter
func (r TaobaoTbkRelationRefundAPIRequest) GetSearchOption() *TopApiRefundRptOption {
	return r._searchOption
}
