package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailElectronicCertificatePreConfirmAPIRequest 贩卖机开始核销接口 API请求
// alibaba.retail.electronic.certificate.pre.confirm
//
// 零售终端贩卖机开始核销接口,返回待领的商品ID
type AlibabaRetailElectronicCertificatePreConfirmAPIRequest struct {
	model.Params
	// 设备ID
	_deviceId string
	// 核销码
	_code int64
}

// NewAlibabaRetailElectronicCertificatePreConfirmRequest 初始化AlibabaRetailElectronicCertificatePreConfirmAPIRequest对象
func NewAlibabaRetailElectronicCertificatePreConfirmRequest() *AlibabaRetailElectronicCertificatePreConfirmAPIRequest {
	return &AlibabaRetailElectronicCertificatePreConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailElectronicCertificatePreConfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.electronic.certificate.pre.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailElectronicCertificatePreConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailElectronicCertificatePreConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备ID
func (r *AlibabaRetailElectronicCertificatePreConfirmAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaRetailElectronicCertificatePreConfirmAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetCode is Code Setter
// 核销码
func (r *AlibabaRetailElectronicCertificatePreConfirmAPIRequest) SetCode(_code int64) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaRetailElectronicCertificatePreConfirmAPIRequest) GetCode() int64 {
	return r._code
}
