package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailelectroniccertificateconfirmAPIRequest 确认核销接口 API请求
// alibaba.retail.electronic.certificate.confirm
//
// 确认核销接口
type AlibabaretailelectroniccertificateconfirmAPIRequest struct {
	model.Params
	// 设备ID
	_deviceId string
	// 核销码
	_code int64
	// 商品ID
	_itemId int64
}

// NewAlibabaretailelectroniccertificateconfirmRequest 初始化AlibabaretailelectroniccertificateconfirmAPIRequest对象
func NewAlibabaretailelectroniccertificateconfirmRequest() *AlibabaretailelectroniccertificateconfirmAPIRequest {
	return &AlibabaretailelectroniccertificateconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailelectroniccertificateconfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.electronic.certificate.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailelectroniccertificateconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailelectroniccertificateconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备ID
func (r *AlibabaretailelectroniccertificateconfirmAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaretailelectroniccertificateconfirmAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetCode is Code Setter
// 核销码
func (r *AlibabaretailelectroniccertificateconfirmAPIRequest) SetCode(_code int64) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaretailelectroniccertificateconfirmAPIRequest) GetCode() int64 {
	return r._code
}

// SetItemId is ItemId Setter
// 商品ID
func (r *AlibabaretailelectroniccertificateconfirmAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaretailelectroniccertificateconfirmAPIRequest) GetItemId() int64 {
	return r._itemId
}
