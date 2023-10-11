package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventoryQueryAPIRequest 库存查询接口（多商品） API请求
// taobao.qimen.inventory.query
//
// taobao.qimen.inventory.query
type TaobaoQimenInventoryQueryAPIRequest struct {
	model.Params
	//
	_request *InventoryQueryRequest
}

// NewTaobaoQimenInventoryQueryRequest 初始化TaobaoQimenInventoryQueryAPIRequest对象
func NewTaobaoQimenInventoryQueryRequest() *TaobaoQimenInventoryQueryAPIRequest {
	return &TaobaoQimenInventoryQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventoryQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventoryQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenInventoryQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenInventoryQueryAPIRequest) SetRequest(_request *InventoryQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenInventoryQueryAPIRequest) GetRequest() *InventoryQueryRequest {
	return r._request
}
