package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeapartmentlineAPIRequest 公寓上下架 API请求
// alibaba.alihouse.newhome.apartment.line
//
// 公寓上下架
type AlibabaalihousenewhomeapartmentlineAPIRequest struct {
	model.Params
	// 外部公寓id
	_outerId string
	// 0-下架 1-上架
	_type *model.File
}

// NewAlibabaalihousenewhomeapartmentlineRequest 初始化AlibabaalihousenewhomeapartmentlineAPIRequest对象
func NewAlibabaalihousenewhomeapartmentlineRequest() *AlibabaalihousenewhomeapartmentlineAPIRequest {
	return &AlibabaalihousenewhomeapartmentlineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeapartmentlineAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.apartment.line"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeapartmentlineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeapartmentlineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 外部公寓id
func (r *AlibabaalihousenewhomeapartmentlineAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaalihousenewhomeapartmentlineAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetType is Type Setter
// 0-下架 1-上架
func (r *AlibabaalihousenewhomeapartmentlineAPIRequest) SetType(_type *model.File) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaalihousenewhomeapartmentlineAPIRequest) GetType() *model.File {
	return r._type
}
