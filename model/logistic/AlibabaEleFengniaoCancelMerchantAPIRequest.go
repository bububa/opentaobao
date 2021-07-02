package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoCancelMerchantAPIRequest 商户取消 API请求
// alibaba.ele.fengniao.cancel.merchant
//
// 商户取消配送
type AlibabaEleFengniaoCancelMerchantAPIRequest struct {
	model.Params
	// 参数param
	_param *Param
}

// NewAlibabaEleFengniaoCancelMerchantRequest 初始化AlibabaEleFengniaoCancelMerchantAPIRequest对象
func NewAlibabaEleFengniaoCancelMerchantRequest() *AlibabaEleFengniaoCancelMerchantAPIRequest {
	return &AlibabaEleFengniaoCancelMerchantAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoCancelMerchantAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.cancel.merchant"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoCancelMerchantAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 参数param
func (r *AlibabaEleFengniaoCancelMerchantAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaEleFengniaoCancelMerchantAPIRequest) GetParam() *Param {
	return r._param
}
