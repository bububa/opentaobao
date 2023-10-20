package logistic

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleFengniaoCancelMerchantAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleFengniaoCancelMerchantAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.fengniao.cancel.merchant"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleFengniaoCancelMerchantAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleFengniaoCancelMerchantAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 参数param
func (r *AlibabaEleFengniaoCancelMerchantAPIRequest) SetParam(_param *Param) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaEleFengniaoCancelMerchantAPIRequest) GetParam() *Param {
	return r._param
}

var poolAlibabaEleFengniaoCancelMerchantAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleFengniaoCancelMerchantRequest()
	},
}

// GetAlibabaEleFengniaoCancelMerchantRequest 从 sync.Pool 获取 AlibabaEleFengniaoCancelMerchantAPIRequest
func GetAlibabaEleFengniaoCancelMerchantAPIRequest() *AlibabaEleFengniaoCancelMerchantAPIRequest {
	return poolAlibabaEleFengniaoCancelMerchantAPIRequest.Get().(*AlibabaEleFengniaoCancelMerchantAPIRequest)
}

// ReleaseAlibabaEleFengniaoCancelMerchantAPIRequest 将 AlibabaEleFengniaoCancelMerchantAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleFengniaoCancelMerchantAPIRequest(v *AlibabaEleFengniaoCancelMerchantAPIRequest) {
	v.Reset()
	poolAlibabaEleFengniaoCancelMerchantAPIRequest.Put(v)
}
