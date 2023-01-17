package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupBigbagContentCancelAPIRequest 进口大包取消 API请求
// cainiao.global.im.pickup.bigbag.content.cancel
//
// 进口大包取消
type CainiaoGlobalImPickupBigbagContentCancelAPIRequest struct {
	model.Params
	// 大包取消请求参数
	_bigbagCancelRequest *BigbagCancelRequest
}

// NewCainiaoGlobalImPickupBigbagContentCancelRequest 初始化CainiaoGlobalImPickupBigbagContentCancelAPIRequest对象
func NewCainiaoGlobalImPickupBigbagContentCancelRequest() *CainiaoGlobalImPickupBigbagContentCancelAPIRequest {
	return &CainiaoGlobalImPickupBigbagContentCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalImPickupBigbagContentCancelAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.bigbag.content.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalImPickupBigbagContentCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalImPickupBigbagContentCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBigbagCancelRequest is BigbagCancelRequest Setter
// 大包取消请求参数
func (r *CainiaoGlobalImPickupBigbagContentCancelAPIRequest) SetBigbagCancelRequest(_bigbagCancelRequest *BigbagCancelRequest) error {
	r._bigbagCancelRequest = _bigbagCancelRequest
	r.Set("bigbag_cancel_request", _bigbagCancelRequest)
	return nil
}

// GetBigbagCancelRequest BigbagCancelRequest Getter
func (r CainiaoGlobalImPickupBigbagContentCancelAPIRequest) GetBigbagCancelRequest() *BigbagCancelRequest {
	return r._bigbagCancelRequest
}
