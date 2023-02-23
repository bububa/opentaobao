package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalHandoverContentSubbagAddAPIRequest 预约单下追加大包 API请求
// cainiao.global.handover.content.subbag.add
//
// 预约单下追加大包
type CainiaoGlobalHandoverContentSubbagAddAPIRequest struct {
	model.Params
	// 预约号(大包LP)
	_handoverContentCode string
	// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
	_client string
	// 多语言
	_locale string
	// 追加大包数量
	_subbagCount int64
}

// NewCainiaoGlobalHandoverContentSubbagAddRequest 初始化CainiaoGlobalHandoverContentSubbagAddAPIRequest对象
func NewCainiaoGlobalHandoverContentSubbagAddRequest() *CainiaoGlobalHandoverContentSubbagAddAPIRequest {
	return &CainiaoGlobalHandoverContentSubbagAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalHandoverContentSubbagAddAPIRequest) GetApiMethodName() string {
	return "cainiao.global.handover.content.subbag.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalHandoverContentSubbagAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoGlobalHandoverContentSubbagAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHandoverContentCode is HandoverContentCode Setter
// 预约号(大包LP)
func (r *CainiaoGlobalHandoverContentSubbagAddAPIRequest) SetHandoverContentCode(_handoverContentCode string) error {
	r._handoverContentCode = _handoverContentCode
	r.Set("handover_content_code", _handoverContentCode)
	return nil
}

// GetHandoverContentCode HandoverContentCode Getter
func (r CainiaoGlobalHandoverContentSubbagAddAPIRequest) GetHandoverContentCode() string {
	return r._handoverContentCode
}

// SetClient is Client Setter
// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverContentSubbagAddAPIRequest) SetClient(_client string) error {
	r._client = _client
	r.Set("client", _client)
	return nil
}

// GetClient Client Getter
func (r CainiaoGlobalHandoverContentSubbagAddAPIRequest) GetClient() string {
	return r._client
}

// SetLocale is Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverContentSubbagAddAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaoGlobalHandoverContentSubbagAddAPIRequest) GetLocale() string {
	return r._locale
}

// SetSubbagCount is SubbagCount Setter
// 追加大包数量
func (r *CainiaoGlobalHandoverContentSubbagAddAPIRequest) SetSubbagCount(_subbagCount int64) error {
	r._subbagCount = _subbagCount
	r.Set("subbag_count", _subbagCount)
	return nil
}

// GetSubbagCount SubbagCount Getter
func (r CainiaoGlobalHandoverContentSubbagAddAPIRequest) GetSubbagCount() int64 {
	return r._subbagCount
}
