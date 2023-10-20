package wlbimports

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupBigbagInfoAPIRequest 大包状态查询 API请求
// cainiao.global.im.pickup.bigbag.info
//
// 大包状态查询
type CainiaoGlobalImPickupBigbagInfoAPIRequest struct {
	model.Params
	// 请求参数
	_bigbagStatusRequest *BigbagStatusRequest
}

// NewCainiaoGlobalImPickupBigbagInfoRequest 初始化CainiaoGlobalImPickupBigbagInfoAPIRequest对象
func NewCainiaoGlobalImPickupBigbagInfoRequest() *CainiaoGlobalImPickupBigbagInfoAPIRequest {
	return &CainiaoGlobalImPickupBigbagInfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoGlobalImPickupBigbagInfoAPIRequest) Reset() {
	r._bigbagStatusRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalImPickupBigbagInfoAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.bigbag.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalImPickupBigbagInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalImPickupBigbagInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBigbagStatusRequest is BigbagStatusRequest Setter
// 请求参数
func (r *CainiaoGlobalImPickupBigbagInfoAPIRequest) SetBigbagStatusRequest(_bigbagStatusRequest *BigbagStatusRequest) error {
	r._bigbagStatusRequest = _bigbagStatusRequest
	r.Set("bigbag_status_request", _bigbagStatusRequest)
	return nil
}

// GetBigbagStatusRequest BigbagStatusRequest Getter
func (r CainiaoGlobalImPickupBigbagInfoAPIRequest) GetBigbagStatusRequest() *BigbagStatusRequest {
	return r._bigbagStatusRequest
}

var poolCainiaoGlobalImPickupBigbagInfoAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoGlobalImPickupBigbagInfoRequest()
	},
}

// GetCainiaoGlobalImPickupBigbagInfoRequest 从 sync.Pool 获取 CainiaoGlobalImPickupBigbagInfoAPIRequest
func GetCainiaoGlobalImPickupBigbagInfoAPIRequest() *CainiaoGlobalImPickupBigbagInfoAPIRequest {
	return poolCainiaoGlobalImPickupBigbagInfoAPIRequest.Get().(*CainiaoGlobalImPickupBigbagInfoAPIRequest)
}

// ReleaseCainiaoGlobalImPickupBigbagInfoAPIRequest 将 CainiaoGlobalImPickupBigbagInfoAPIRequest 放入 sync.Pool
func ReleaseCainiaoGlobalImPickupBigbagInfoAPIRequest(v *CainiaoGlobalImPickupBigbagInfoAPIRequest) {
	v.Reset()
	poolCainiaoGlobalImPickupBigbagInfoAPIRequest.Put(v)
}
