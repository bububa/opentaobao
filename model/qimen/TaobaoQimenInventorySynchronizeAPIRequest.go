package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimeninventorysynchronizeAPIRequest 库存状态同步接口 API请求
// taobao.qimen.inventory.synchronize
//
// ERP通过该接口同步指定商品的库存信息
type TaobaoqimeninventorysynchronizeAPIRequest struct {
	model.Params
	//
	_request *TaobaoqimeninventorysynchronizeRequest
}

// NewTaobaoqimeninventorysynchronizeRequest 初始化TaobaoqimeninventorysynchronizeAPIRequest对象
func NewTaobaoqimeninventorysynchronizeRequest() *TaobaoqimeninventorysynchronizeAPIRequest {
	return &TaobaoqimeninventorysynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimeninventorysynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventory.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimeninventorysynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimeninventorysynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimeninventorysynchronizeAPIRequest) SetRequest(_request *TaobaoqimeninventorysynchronizeRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimeninventorysynchronizeAPIRequest) GetRequest() *TaobaoqimeninventorysynchronizeRequest {
	return r._request
}
