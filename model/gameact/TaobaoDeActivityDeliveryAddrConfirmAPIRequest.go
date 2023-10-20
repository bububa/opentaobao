package gameact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodeactivitydeliveryaddrconfirmAPIRequest 用户收件地址确认 API请求
// taobao.de.activity.delivery.addr.confirm
//
// 用户收件地址确认
type TaobaodeactivitydeliveryaddrconfirmAPIRequest struct {
	model.Params
	// 加密流水号
	_serialNumber string
	// 地址Sign
	_addressSign string
}

// NewTaobaodeactivitydeliveryaddrconfirmRequest 初始化TaobaodeactivitydeliveryaddrconfirmAPIRequest对象
func NewTaobaodeactivitydeliveryaddrconfirmRequest() *TaobaodeactivitydeliveryaddrconfirmAPIRequest {
	return &TaobaodeactivitydeliveryaddrconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodeactivitydeliveryaddrconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.de.activity.delivery.addr.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodeactivitydeliveryaddrconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodeactivitydeliveryaddrconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSerialNumber is SerialNumber Setter
// 加密流水号
func (r *TaobaodeactivitydeliveryaddrconfirmAPIRequest) SetSerialNumber(_serialNumber string) error {
	r._serialNumber = _serialNumber
	r.Set("serial_number", _serialNumber)
	return nil
}

// GetSerialNumber SerialNumber Getter
func (r TaobaodeactivitydeliveryaddrconfirmAPIRequest) GetSerialNumber() string {
	return r._serialNumber
}

// SetAddressSign is AddressSign Setter
// 地址Sign
func (r *TaobaodeactivitydeliveryaddrconfirmAPIRequest) SetAddressSign(_addressSign string) error {
	r._addressSign = _addressSign
	r.Set("address_sign", _addressSign)
	return nil
}

// GetAddressSign AddressSign Getter
func (r TaobaodeactivitydeliveryaddrconfirmAPIRequest) GetAddressSign() string {
	return r._addressSign
}
