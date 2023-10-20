package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeapartmentouteridAPIRequest 公寓更新outerid API请求
// alibaba.alihouse.newhome.apartment.outerid
//
// 公寓更新outerid
type AlibabaalihousenewhomeapartmentouteridAPIRequest struct {
	model.Params
	// 公寓outerid
	_outerId string
	// 公寓ecode
	_eCode string
}

// NewAlibabaalihousenewhomeapartmentouteridRequest 初始化AlibabaalihousenewhomeapartmentouteridAPIRequest对象
func NewAlibabaalihousenewhomeapartmentouteridRequest() *AlibabaalihousenewhomeapartmentouteridAPIRequest {
	return &AlibabaalihousenewhomeapartmentouteridAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeapartmentouteridAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.apartment.outerid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeapartmentouteridAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeapartmentouteridAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 公寓outerid
func (r *AlibabaalihousenewhomeapartmentouteridAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaalihousenewhomeapartmentouteridAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetECode is ECode Setter
// 公寓ecode
func (r *AlibabaalihousenewhomeapartmentouteridAPIRequest) SetECode(_eCode string) error {
	r._eCode = _eCode
	r.Set("e_code", _eCode)
	return nil
}

// GetECode ECode Getter
func (r AlibabaalihousenewhomeapartmentouteridAPIRequest) GetECode() string {
	return r._eCode
}
