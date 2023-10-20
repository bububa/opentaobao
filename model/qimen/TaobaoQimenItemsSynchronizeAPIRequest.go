package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenitemssynchronizeAPIRequest 商品同步接口 (批量) API请求
// taobao.qimen.items.synchronize
//
// ERP调用奇门的接口,批量同步商品信息给WMS
type TaobaoqimenitemssynchronizeAPIRequest struct {
	model.Params
	//
	_request *ItemsSynchronizeRequest
}

// NewTaobaoqimenitemssynchronizeRequest 初始化TaobaoqimenitemssynchronizeAPIRequest对象
func NewTaobaoqimenitemssynchronizeRequest() *TaobaoqimenitemssynchronizeAPIRequest {
	return &TaobaoqimenitemssynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenitemssynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.items.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenitemssynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenitemssynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenitemssynchronizeAPIRequest) SetRequest(_request *ItemsSynchronizeRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenitemssynchronizeAPIRequest) GetRequest() *ItemsSynchronizeRequest {
	return r._request
}
