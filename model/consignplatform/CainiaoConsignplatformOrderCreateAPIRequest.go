package consignplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoconsignplatformordercreateAPIRequest 菜鸟发货工作台创建订单 API请求
// cainiao.consignplatform.order.create
//
// 菜鸟发货工作台，商家或者isv通过api进行订单写入操作
type CainiaoconsignplatformordercreateAPIRequest struct {
	model.Params
	// 订单创建入参
	_createRequest *OrderCreateRequest
}

// NewCainiaoconsignplatformordercreateRequest 初始化CainiaoconsignplatformordercreateAPIRequest对象
func NewCainiaoconsignplatformordercreateRequest() *CainiaoconsignplatformordercreateAPIRequest {
	return &CainiaoconsignplatformordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoconsignplatformordercreateAPIRequest) GetApiMethodName() string {
	return "cainiao.consignplatform.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoconsignplatformordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoconsignplatformordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateRequest is CreateRequest Setter
// 订单创建入参
func (r *CainiaoconsignplatformordercreateAPIRequest) SetCreateRequest(_createRequest *OrderCreateRequest) error {
	r._createRequest = _createRequest
	r.Set("create_request", _createRequest)
	return nil
}

// GetCreateRequest CreateRequest Getter
func (r CainiaoconsignplatformordercreateAPIRequest) GetCreateRequest() *OrderCreateRequest {
	return r._createRequest
}
