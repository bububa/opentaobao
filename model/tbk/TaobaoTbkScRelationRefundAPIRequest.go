package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScRelationRefundAPIRequest 淘宝客-服务商-维权退款订单查询 API请求
// taobao.tbk.sc.relation.refund
//
// 淘宝客维权退款订单查询-渠道管理和会员运营管理专用
type TaobaoTbkScRelationRefundAPIRequest struct {
	model.Params
	// 参数option
	_searchOption *TopApiRefundRptOption
}

// NewTaobaoTbkScRelationRefundRequest 初始化TaobaoTbkScRelationRefundAPIRequest对象
func NewTaobaoTbkScRelationRefundRequest() *TaobaoTbkScRelationRefundAPIRequest {
	return &TaobaoTbkScRelationRefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScRelationRefundAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.relation.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScRelationRefundAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSearchOption is SearchOption Setter
// 参数option
func (r *TaobaoTbkScRelationRefundAPIRequest) SetSearchOption(_searchOption *TopApiRefundRptOption) error {
	r._searchOption = _searchOption
	r.Set("search_option", _searchOption)
	return nil
}

// GetSearchOption SearchOption Getter
func (r TaobaoTbkScRelationRefundAPIRequest) GetSearchOption() *TopApiRefundRptOption {
	return r._searchOption
}
