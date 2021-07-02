package consignplatform

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoConsignplatformOrderCreateAPIRequest) GetApiMethodName() string {
	return "cainiao.consignplatform.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoConsignplatformOrderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
