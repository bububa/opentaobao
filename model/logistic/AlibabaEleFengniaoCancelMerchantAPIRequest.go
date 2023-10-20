package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaelefengniaocancelmerchantAPIRequest 商户取消 API请求
// alibaba.ele.fengniao.cancel.merchant
//
// 商户取消配送
type AlibabaelefengniaocancelmerchantAPIRequest struct {
	model.Params
	// 参数param
	_param *Param
}

// NewAlibabaelefengniaocancelmerchantRequest 初始化AlibabaelefengniaocancelmerchantAPIRequest对象
func NewAlibabaelefengniaocancelmerchantRequest() *AlibabaelefengniaocancelmerchantAPIRequest {
	return &AlibabaelefengniaocancelmerchantAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaelefengniaocancelmerchantAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.cancel.merchant"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaelefengniaocancelmerchantAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaelefengniaocancelmerchantAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 参数param
func (r *AlibabaelefengniaocancelmerchantAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaelefengniaocancelmerchantAPIRequest) GetParam() *Param {
	return r._param
}
