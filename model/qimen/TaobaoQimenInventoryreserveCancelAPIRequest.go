package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenInventoryreserveCancelAPIRequest
库存预占取消接口 API请求
taobao.qimen.inventoryreserve.cancel

库存预占取消 */
type TaobaoQimenInventoryreserveCancelAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenInventoryreserveCancelRequest
}

// NewTaobaoQimenInventoryreserveCancelRequest 初始化TaobaoQimenInventoryreserveCancelAPIRequest对象
func NewTaobaoQimenInventoryreserveCancelRequest() *TaobaoQimenInventoryreserveCancelAPIRequest {
	return &TaobaoQimenInventoryreserveCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventoryreserveCancelAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventoryreserve.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventoryreserveCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
//
func (r *TaobaoQimenInventoryreserveCancelAPIRequest) SetRequest(_request *TaobaoQimenInventoryreserveCancelRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenInventoryreserveCancelAPIRequest) GetRequest() *TaobaoQimenInventoryreserveCancelRequest {
	return r._request
}
