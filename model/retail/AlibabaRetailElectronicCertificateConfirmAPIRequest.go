package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailElectronicCertificateConfirmAPIRequest
确认核销接口 API请求
alibaba.retail.electronic.certificate.confirm

确认核销接口 */
type AlibabaRetailElectronicCertificateConfirmAPIRequest struct {
	model.Params
	// 核销码
	_code int64
	// 商品ID
	_itemId int64
	// 设备ID
	_deviceId string
}

// NewAlibabaRetailElectronicCertificateConfirmRequest 初始化AlibabaRetailElectronicCertificateConfirmAPIRequest对象
func NewAlibabaRetailElectronicCertificateConfirmRequest() *AlibabaRetailElectronicCertificateConfirmAPIRequest {
	return &AlibabaRetailElectronicCertificateConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailElectronicCertificateConfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.electronic.certificate.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailElectronicCertificateConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Code Setter
// 核销码
func (r *AlibabaRetailElectronicCertificateConfirmAPIRequest) SetCode(_code int64) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// Get Code Getter
func (r AlibabaRetailElectronicCertificateConfirmAPIRequest) GetCode() int64 {
	return r._code
}

// Set is ItemId Setter
// 商品ID
func (r *AlibabaRetailElectronicCertificateConfirmAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r AlibabaRetailElectronicCertificateConfirmAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is DeviceId Setter
// 设备ID
func (r *AlibabaRetailElectronicCertificateConfirmAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// Get DeviceId Getter
func (r AlibabaRetailElectronicCertificateConfirmAPIRequest) GetDeviceId() string {
	return r._deviceId
}
