package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantchannelorderprecheckAPIRequest 前置校验商品是否可下单作业 API请求
// alibaba.tcls.aelophy.merchant.channel.order.precheck
//
// 前置校验商品是否可下单作业
type AlibabatclsaelophymerchantchannelorderprecheckAPIRequest struct {
	model.Params
	// 前置校验入参
	_preCheckRequest *PreCheckRequest
}

// NewAlibabatclsaelophymerchantchannelorderprecheckRequest 初始化AlibabatclsaelophymerchantchannelorderprecheckAPIRequest对象
func NewAlibabatclsaelophymerchantchannelorderprecheckRequest() *AlibabatclsaelophymerchantchannelorderprecheckAPIRequest {
	return &AlibabatclsaelophymerchantchannelorderprecheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophymerchantchannelorderprecheckAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.order.precheck"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophymerchantchannelorderprecheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophymerchantchannelorderprecheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPreCheckRequest is PreCheckRequest Setter
// 前置校验入参
func (r *AlibabatclsaelophymerchantchannelorderprecheckAPIRequest) SetPreCheckRequest(_preCheckRequest *PreCheckRequest) error {
	r._preCheckRequest = _preCheckRequest
	r.Set("pre_check_request", _preCheckRequest)
	return nil
}

// GetPreCheckRequest PreCheckRequest Getter
func (r AlibabatclsaelophymerchantchannelorderprecheckAPIRequest) GetPreCheckRequest() *PreCheckRequest {
	return r._preCheckRequest
}
