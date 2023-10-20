package newretail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitapaddressgetAPIRequest getApAddressByMacNew API请求
// alibaba.it.ap.address.get
//
// 根据ap 的mac地址查询ap的结构化位置信息
type AlibabaitapaddressgetAPIRequest struct {
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

// NewAlibabaitapaddressgetRequest 初始化AlibabaitapaddressgetAPIRequest对象
func NewAlibabaitapaddressgetRequest() *AlibabaitapaddressgetAPIRequest {
	return &AlibabaitapaddressgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaitapaddressgetAPIRequest) GetApiMethodName() string {
	return "alibaba.it.ap.address.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaitapaddressgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaitapaddressgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSignature is Signature Setter
// 签名
func (r *AlibabaitapaddressgetAPIRequest) SetSignature(_signature string) error {
	r._signature = _signature
	r.Set("signature", _signature)
	return nil
}

// GetSignature Signature Getter
func (r AlibabaitapaddressgetAPIRequest) GetSignature() string {
	return r._signature
}

// SetLanguage is Language Setter
// 语言
func (r *AlibabaitapaddressgetAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AlibabaitapaddressgetAPIRequest) GetLanguage() string {
	return r._language
}

// SetAppKeyInternal is AppKeyInternal Setter
// 分配的ak
func (r *AlibabaitapaddressgetAPIRequest) SetAppKeyInternal(_appKeyInternal string) error {
	r._appKeyInternal = _appKeyInternal
	r.Set("app_key_internal", _appKeyInternal)
	return nil
}

// GetAppKeyInternal AppKeyInternal Getter
func (r AlibabaitapaddressgetAPIRequest) GetAppKeyInternal() string {
	return r._appKeyInternal
}

// SetMac is Mac Setter
// ap的mac
func (r *AlibabaitapaddressgetAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// GetMac Mac Getter
func (r AlibabaitapaddressgetAPIRequest) GetMac() string {
	return r._mac
}

// SetTimestampInternal is TimestampInternal Setter
// 当前时间戳，毫秒
func (r *AlibabaitapaddressgetAPIRequest) SetTimestampInternal(_timestampInternal int64) error {
	r._timestampInternal = _timestampInternal
	r.Set("timestamp_internal", _timestampInternal)
	return nil
}

// GetTimestampInternal TimestampInternal Getter
func (r AlibabaitapaddressgetAPIRequest) GetTimestampInternal() int64 {
	return r._timestampInternal
}
