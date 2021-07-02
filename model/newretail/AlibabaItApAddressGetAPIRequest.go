package newretail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItApAddressGetAPIRequest getApAddressByMacNew API请求
// alibaba.it.ap.address.get
//
// 根据ap 的mac地址查询ap的结构化位置信息
type AlibabaItApAddressGetAPIRequest struct {
	model.Params
	// 签名
	_signature string
	// 语言
	_language string
	// 分配的ak
	_appKeyInternal string
	// ap的mac
	_mac string
	// 当前时间戳，毫秒
	_timestampInternal int64
}

// NewAlibabaItApAddressGetRequest 初始化AlibabaItApAddressGetAPIRequest对象
func NewAlibabaItApAddressGetRequest() *AlibabaItApAddressGetAPIRequest {
	return &AlibabaItApAddressGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItApAddressGetAPIRequest) GetApiMethodName() string {
	return "alibaba.it.ap.address.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItApAddressGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSignature is Signature Setter
// 签名
func (r *AlibabaItApAddressGetAPIRequest) SetSignature(_signature string) error {
	r._signature = _signature
	r.Set("signature", _signature)
	return nil
}

// GetSignature Signature Getter
func (r AlibabaItApAddressGetAPIRequest) GetSignature() string {
	return r._signature
}

// SetLanguage is Language Setter
// 语言
func (r *AlibabaItApAddressGetAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AlibabaItApAddressGetAPIRequest) GetLanguage() string {
	return r._language
}

// SetAppKeyInternal is AppKeyInternal Setter
// 分配的ak
func (r *AlibabaItApAddressGetAPIRequest) SetAppKeyInternal(_appKeyInternal string) error {
	r._appKeyInternal = _appKeyInternal
	r.Set("app_key_internal", _appKeyInternal)
	return nil
}

// GetAppKeyInternal AppKeyInternal Getter
func (r AlibabaItApAddressGetAPIRequest) GetAppKeyInternal() string {
	return r._appKeyInternal
}

// SetMac is Mac Setter
// ap的mac
func (r *AlibabaItApAddressGetAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// GetMac Mac Getter
func (r AlibabaItApAddressGetAPIRequest) GetMac() string {
	return r._mac
}

// SetTimestampInternal is TimestampInternal Setter
// 当前时间戳，毫秒
func (r *AlibabaItApAddressGetAPIRequest) SetTimestampInternal(_timestampInternal int64) error {
	r._timestampInternal = _timestampInternal
	r.Set("timestamp_internal", _timestampInternal)
	return nil
}

// GetTimestampInternal TimestampInternal Getter
func (r AlibabaItApAddressGetAPIRequest) GetTimestampInternal() int64 {
	return r._timestampInternal
}
