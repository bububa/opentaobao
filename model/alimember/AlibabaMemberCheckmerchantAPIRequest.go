package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamembercheckmerchantAPIRequest 校验商家身份 API请求
// alibaba.member.checkmerchant
//
// 校验商家身份
type AlibabamembercheckmerchantAPIRequest struct {
	model.Params
	// 混淆后的商家id
	_openMerchantId string
}

// NewAlibabamembercheckmerchantRequest 初始化AlibabamembercheckmerchantAPIRequest对象
func NewAlibabamembercheckmerchantRequest() *AlibabamembercheckmerchantAPIRequest {
	return &AlibabamembercheckmerchantAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamembercheckmerchantAPIRequest) GetApiMethodName() string {
	return "alibaba.member.checkmerchant"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamembercheckmerchantAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamembercheckmerchantAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenMerchantId is OpenMerchantId Setter
// 混淆后的商家id
func (r *AlibabamembercheckmerchantAPIRequest) SetOpenMerchantId(_openMerchantId string) error {
	r._openMerchantId = _openMerchantId
	r.Set("open_merchant_id", _openMerchantId)
	return nil
}

// GetOpenMerchantId OpenMerchantId Getter
func (r AlibabamembercheckmerchantAPIRequest) GetOpenMerchantId() string {
	return r._openMerchantId
}
