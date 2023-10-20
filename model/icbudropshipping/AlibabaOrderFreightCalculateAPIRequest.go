package icbudropshipping

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOrderFreightCalculateAPIRequest 阿里巴巴下单场景运费方案计算 API请求
// alibaba.order.freight.calculate
//
// icbu开展 drop shipping 业务，阿里巴巴下单场景运费方案计算
// alibaba Create order scenario freight calculation
type AlibabaOrderFreightCalculateAPIRequest struct {
	model.Params
	// {}
	_paramMultiFreightTemplateRequest *MultiFreightTemplateRequest
}

// NewAlibabaOrderFreightCalculateRequest 初始化AlibabaOrderFreightCalculateAPIRequest对象
func NewAlibabaOrderFreightCalculateRequest() *AlibabaOrderFreightCalculateAPIRequest {
	return &AlibabaOrderFreightCalculateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaOrderFreightCalculateAPIRequest) Reset() {
	r._paramMultiFreightTemplateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOrderFreightCalculateAPIRequest) GetApiMethodName() string {
	return "alibaba.order.freight.calculate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOrderFreightCalculateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaOrderFreightCalculateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamMultiFreightTemplateRequest is ParamMultiFreightTemplateRequest Setter
// {}
func (r *AlibabaOrderFreightCalculateAPIRequest) SetParamMultiFreightTemplateRequest(_paramMultiFreightTemplateRequest *MultiFreightTemplateRequest) error {
	r._paramMultiFreightTemplateRequest = _paramMultiFreightTemplateRequest
	r.Set("param_multi_freight_template_request", _paramMultiFreightTemplateRequest)
	return nil
}

// GetParamMultiFreightTemplateRequest ParamMultiFreightTemplateRequest Getter
func (r AlibabaOrderFreightCalculateAPIRequest) GetParamMultiFreightTemplateRequest() *MultiFreightTemplateRequest {
	return r._paramMultiFreightTemplateRequest
}

var poolAlibabaOrderFreightCalculateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaOrderFreightCalculateRequest()
	},
}

// GetAlibabaOrderFreightCalculateRequest 从 sync.Pool 获取 AlibabaOrderFreightCalculateAPIRequest
func GetAlibabaOrderFreightCalculateAPIRequest() *AlibabaOrderFreightCalculateAPIRequest {
	return poolAlibabaOrderFreightCalculateAPIRequest.Get().(*AlibabaOrderFreightCalculateAPIRequest)
}

// ReleaseAlibabaOrderFreightCalculateAPIRequest 将 AlibabaOrderFreightCalculateAPIRequest 放入 sync.Pool
func ReleaseAlibabaOrderFreightCalculateAPIRequest(v *AlibabaOrderFreightCalculateAPIRequest) {
	v.Reset()
	poolAlibabaOrderFreightCalculateAPIRequest.Put(v)
}
