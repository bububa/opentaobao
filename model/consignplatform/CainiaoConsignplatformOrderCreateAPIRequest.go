package consignplatform

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoConsignplatformOrderCreateAPIRequest 菜鸟发货工作台创建订单 API请求
// cainiao.consignplatform.order.create
//
// 菜鸟发货工作台，商家或者isv通过api进行订单写入操作
type CainiaoConsignplatformOrderCreateAPIRequest struct {
	model.Params
	// 订单创建入参
	_createRequest *OrderCreateRequest
}

// NewCainiaoConsignplatformOrderCreateRequest 初始化CainiaoConsignplatformOrderCreateAPIRequest对象
func NewCainiaoConsignplatformOrderCreateRequest() *CainiaoConsignplatformOrderCreateAPIRequest {
	return &CainiaoConsignplatformOrderCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoConsignplatformOrderCreateAPIRequest) Reset() {
	r._createRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoConsignplatformOrderCreateAPIRequest) GetApiMethodName() string {
	return "cainiao.consignplatform.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoConsignplatformOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoConsignplatformOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateRequest is CreateRequest Setter
// 订单创建入参
func (r *CainiaoConsignplatformOrderCreateAPIRequest) SetCreateRequest(_createRequest *OrderCreateRequest) error {
	r._createRequest = _createRequest
	r.Set("create_request", _createRequest)
	return nil
}

// GetCreateRequest CreateRequest Getter
func (r CainiaoConsignplatformOrderCreateAPIRequest) GetCreateRequest() *OrderCreateRequest {
	return r._createRequest
}

var poolCainiaoConsignplatformOrderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoConsignplatformOrderCreateRequest()
	},
}

// GetCainiaoConsignplatformOrderCreateRequest 从 sync.Pool 获取 CainiaoConsignplatformOrderCreateAPIRequest
func GetCainiaoConsignplatformOrderCreateAPIRequest() *CainiaoConsignplatformOrderCreateAPIRequest {
	return poolCainiaoConsignplatformOrderCreateAPIRequest.Get().(*CainiaoConsignplatformOrderCreateAPIRequest)
}

// ReleaseCainiaoConsignplatformOrderCreateAPIRequest 将 CainiaoConsignplatformOrderCreateAPIRequest 放入 sync.Pool
func ReleaseCainiaoConsignplatformOrderCreateAPIRequest(v *CainiaoConsignplatformOrderCreateAPIRequest) {
	v.Reset()
	poolCainiaoConsignplatformOrderCreateAPIRequest.Put(v)
}
