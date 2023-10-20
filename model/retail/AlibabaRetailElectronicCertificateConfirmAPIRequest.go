package retail

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailElectronicCertificateConfirmAPIRequest 确认核销接口 API请求
// alibaba.retail.electronic.certificate.confirm
//
// 确认核销接口
type AlibabaRetailElectronicCertificateConfirmAPIRequest struct {
	model.Params
	// 设备ID
	_deviceId string
	// 核销码
	_code int64
	// 商品ID
	_itemId int64
}

// NewAlibabaRetailElectronicCertificateConfirmRequest 初始化AlibabaRetailElectronicCertificateConfirmAPIRequest对象
func NewAlibabaRetailElectronicCertificateConfirmRequest() *AlibabaRetailElectronicCertificateConfirmAPIRequest {
	return &AlibabaRetailElectronicCertificateConfirmAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailElectronicCertificateConfirmAPIRequest) Reset() {
	r._deviceId = ""
	r._code = 0
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailElectronicCertificateConfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.electronic.certificate.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailElectronicCertificateConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailElectronicCertificateConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备ID
func (r *AlibabaRetailElectronicCertificateConfirmAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaRetailElectronicCertificateConfirmAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetCode is Code Setter
// 核销码
func (r *AlibabaRetailElectronicCertificateConfirmAPIRequest) SetCode(_code int64) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaRetailElectronicCertificateConfirmAPIRequest) GetCode() int64 {
	return r._code
}

// SetItemId is ItemId Setter
// 商品ID
func (r *AlibabaRetailElectronicCertificateConfirmAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaRetailElectronicCertificateConfirmAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolAlibabaRetailElectronicCertificateConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailElectronicCertificateConfirmRequest()
	},
}

// GetAlibabaRetailElectronicCertificateConfirmRequest 从 sync.Pool 获取 AlibabaRetailElectronicCertificateConfirmAPIRequest
func GetAlibabaRetailElectronicCertificateConfirmAPIRequest() *AlibabaRetailElectronicCertificateConfirmAPIRequest {
	return poolAlibabaRetailElectronicCertificateConfirmAPIRequest.Get().(*AlibabaRetailElectronicCertificateConfirmAPIRequest)
}

// ReleaseAlibabaRetailElectronicCertificateConfirmAPIRequest 将 AlibabaRetailElectronicCertificateConfirmAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailElectronicCertificateConfirmAPIRequest(v *AlibabaRetailElectronicCertificateConfirmAPIRequest) {
	v.Reset()
	poolAlibabaRetailElectronicCertificateConfirmAPIRequest.Put(v)
}
