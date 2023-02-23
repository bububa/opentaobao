package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupBigbagContentCreateAPIRequest 大包创建 API请求
// cainiao.global.im.pickup.bigbag.content.create
//
// 大包创建
type CainiaoGlobalImPickupBigbagContentCreateAPIRequest struct {
	model.Params
	// 大包创建入参
	_bigbagCreateRequest *BigbagCreateRequest
}

// NewCainiaoGlobalImPickupBigbagContentCreateRequest 初始化CainiaoGlobalImPickupBigbagContentCreateAPIRequest对象
func NewCainiaoGlobalImPickupBigbagContentCreateRequest() *CainiaoGlobalImPickupBigbagContentCreateAPIRequest {
	return &CainiaoGlobalImPickupBigbagContentCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalImPickupBigbagContentCreateAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.bigbag.content.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalImPickupBigbagContentCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalImPickupBigbagContentCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBigbagCreateRequest is BigbagCreateRequest Setter
// 大包创建入参
func (r *CainiaoGlobalImPickupBigbagContentCreateAPIRequest) SetBigbagCreateRequest(_bigbagCreateRequest *BigbagCreateRequest) error {
	r._bigbagCreateRequest = _bigbagCreateRequest
	r.Set("bigbag_create_request", _bigbagCreateRequest)
	return nil
}

// GetBigbagCreateRequest BigbagCreateRequest Getter
func (r CainiaoGlobalImPickupBigbagContentCreateAPIRequest) GetBigbagCreateRequest() *BigbagCreateRequest {
	return r._bigbagCreateRequest
}
