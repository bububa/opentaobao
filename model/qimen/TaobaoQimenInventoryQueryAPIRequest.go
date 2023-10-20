package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimeninventoryqueryAPIRequest 库存查询接口（多商品） API请求
// taobao.qimen.inventory.query
//
// taobao.qimen.inventory.query
type TaobaoqimeninventoryqueryAPIRequest struct {
	model.Params
	//
	_request *InventoryQueryRequest
}

// NewTaobaoqimeninventoryqueryRequest 初始化TaobaoqimeninventoryqueryAPIRequest对象
func NewTaobaoqimeninventoryqueryRequest() *TaobaoqimeninventoryqueryAPIRequest {
	return &TaobaoqimeninventoryqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimeninventoryqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimeninventoryqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimeninventoryqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimeninventoryqueryAPIRequest) SetRequest(_request *InventoryQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimeninventoryqueryAPIRequest) GetRequest() *InventoryQueryRequest {
	return r._request
}
