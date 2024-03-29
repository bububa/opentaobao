package wlbimports

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest 大包面单查询 API请求
// cainiao.global.im.pickup.bigbag.waybill.info
//
// 大包面单查询
type CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest struct {
	model.Params
	// bigbagId和appointmentOrderId必填一个
	_bigbagWaybillRequest *BigbagWaybillRequest
}

// NewCainiaoGlobalImPickupBigbagWaybillInfoRequest 初始化CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest对象
func NewCainiaoGlobalImPickupBigbagWaybillInfoRequest() *CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest {
	return &CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest) Reset() {
	r._bigbagWaybillRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.bigbag.waybill.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBigbagWaybillRequest is BigbagWaybillRequest Setter
// bigbagId和appointmentOrderId必填一个
func (r *CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest) SetBigbagWaybillRequest(_bigbagWaybillRequest *BigbagWaybillRequest) error {
	r._bigbagWaybillRequest = _bigbagWaybillRequest
	r.Set("bigbag_waybill_request", _bigbagWaybillRequest)
	return nil
}

// GetBigbagWaybillRequest BigbagWaybillRequest Getter
func (r CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest) GetBigbagWaybillRequest() *BigbagWaybillRequest {
	return r._bigbagWaybillRequest
}

var poolCainiaoGlobalImPickupBigbagWaybillInfoAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoGlobalImPickupBigbagWaybillInfoRequest()
	},
}

// GetCainiaoGlobalImPickupBigbagWaybillInfoRequest 从 sync.Pool 获取 CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest
func GetCainiaoGlobalImPickupBigbagWaybillInfoAPIRequest() *CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest {
	return poolCainiaoGlobalImPickupBigbagWaybillInfoAPIRequest.Get().(*CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest)
}

// ReleaseCainiaoGlobalImPickupBigbagWaybillInfoAPIRequest 将 CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest 放入 sync.Pool
func ReleaseCainiaoGlobalImPickupBigbagWaybillInfoAPIRequest(v *CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest) {
	v.Reset()
	poolCainiaoGlobalImPickupBigbagWaybillInfoAPIRequest.Put(v)
}
