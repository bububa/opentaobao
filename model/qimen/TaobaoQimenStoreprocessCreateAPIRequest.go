package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStoreprocessCreateAPIRequest 仓内加工单创建接口 API请求
// taobao.qimen.storeprocess.create
//
// ERP调用奇门的接口,创建仓内加工单
type TaobaoQimenStoreprocessCreateAPIRequest struct {
	model.Params
	//
	_request *StoreProcessCreateRequest
}

// NewTaobaoQimenStoreprocessCreateRequest 初始化TaobaoQimenStoreprocessCreateAPIRequest对象
func NewTaobaoQimenStoreprocessCreateRequest() *TaobaoQimenStoreprocessCreateAPIRequest {
	return &TaobaoQimenStoreprocessCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenStoreprocessCreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.storeprocess.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenStoreprocessCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenStoreprocessCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenStoreprocessCreateAPIRequest) SetRequest(_request *StoreProcessCreateRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenStoreprocessCreateAPIRequest) GetRequest() *StoreProcessCreateRequest {
	return r._request
}
