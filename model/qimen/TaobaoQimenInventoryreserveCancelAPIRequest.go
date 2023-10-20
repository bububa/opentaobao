package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimeninventoryreservecancelAPIRequest 库存预占取消接口 API请求
// taobao.qimen.inventoryreserve.cancel
//
// 库存预占取消
type TaobaoqimeninventoryreservecancelAPIRequest struct {
	model.Params
	//
	_request *TaobaoqimeninventoryreservecancelRequest
}

// NewTaobaoqimeninventoryreservecancelRequest 初始化TaobaoqimeninventoryreservecancelAPIRequest对象
func NewTaobaoqimeninventoryreservecancelRequest() *TaobaoqimeninventoryreservecancelAPIRequest {
	return &TaobaoqimeninventoryreservecancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimeninventoryreservecancelAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventoryreserve.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimeninventoryreservecancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimeninventoryreservecancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimeninventoryreservecancelAPIRequest) SetRequest(_request *TaobaoqimeninventoryreservecancelRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimeninventoryreservecancelAPIRequest) GetRequest() *TaobaoqimeninventoryreservecancelRequest {
	return r._request
}
