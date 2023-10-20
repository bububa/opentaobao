package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimeninventorybatchqueryAPIRequest 商品单仓批次库存查询接口 API请求
// taobao.qimen.inventorybatch.query
//
// ERP 通过该接口查询指定商品的单仓批次库存
type TaobaoqimeninventorybatchqueryAPIRequest struct {
	model.Params
	// request
	_request *TaobaoqimeninventorybatchqueryRequest
}

// NewTaobaoqimeninventorybatchqueryRequest 初始化TaobaoqimeninventorybatchqueryAPIRequest对象
func NewTaobaoqimeninventorybatchqueryRequest() *TaobaoqimeninventorybatchqueryAPIRequest {
	return &TaobaoqimeninventorybatchqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimeninventorybatchqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventorybatch.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimeninventorybatchqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimeninventorybatchqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
// request
func (r *TaobaoqimeninventorybatchqueryAPIRequest) SetRequest(_request *TaobaoqimeninventorybatchqueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimeninventorybatchqueryAPIRequest) GetRequest() *TaobaoqimeninventorybatchqueryRequest {
	return r._request
}
