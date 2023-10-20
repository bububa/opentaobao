package icbu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductSchemaAddDraftAPIRequest （新）ICBU商品发布草稿 API请求
// alibaba.icbu.product.schema.add.draft
//
// 提供发布ICBU商品草稿的入口
type AlibabaIcbuProductSchemaAddDraftAPIRequest struct {
	model.Params
	// 发布入参
	_paramProductTopPublishRequest *ProductTopPublishRequest
}

// NewAlibabaIcbuProductSchemaAddDraftRequest 初始化AlibabaIcbuProductSchemaAddDraftAPIRequest对象
func NewAlibabaIcbuProductSchemaAddDraftRequest() *AlibabaIcbuProductSchemaAddDraftAPIRequest {
	return &AlibabaIcbuProductSchemaAddDraftAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuProductSchemaAddDraftAPIRequest) Reset() {
	r._paramProductTopPublishRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductSchemaAddDraftAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.schema.add.draft"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductSchemaAddDraftAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuProductSchemaAddDraftAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamProductTopPublishRequest is ParamProductTopPublishRequest Setter
// 发布入参
func (r *AlibabaIcbuProductSchemaAddDraftAPIRequest) SetParamProductTopPublishRequest(_paramProductTopPublishRequest *ProductTopPublishRequest) error {
	r._paramProductTopPublishRequest = _paramProductTopPublishRequest
	r.Set("param_product_top_publish_request", _paramProductTopPublishRequest)
	return nil
}

// GetParamProductTopPublishRequest ParamProductTopPublishRequest Getter
func (r AlibabaIcbuProductSchemaAddDraftAPIRequest) GetParamProductTopPublishRequest() *ProductTopPublishRequest {
	return r._paramProductTopPublishRequest
}

var poolAlibabaIcbuProductSchemaAddDraftAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuProductSchemaAddDraftRequest()
	},
}

// GetAlibabaIcbuProductSchemaAddDraftRequest 从 sync.Pool 获取 AlibabaIcbuProductSchemaAddDraftAPIRequest
func GetAlibabaIcbuProductSchemaAddDraftAPIRequest() *AlibabaIcbuProductSchemaAddDraftAPIRequest {
	return poolAlibabaIcbuProductSchemaAddDraftAPIRequest.Get().(*AlibabaIcbuProductSchemaAddDraftAPIRequest)
}

// ReleaseAlibabaIcbuProductSchemaAddDraftAPIRequest 将 AlibabaIcbuProductSchemaAddDraftAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuProductSchemaAddDraftAPIRequest(v *AlibabaIcbuProductSchemaAddDraftAPIRequest) {
	v.Reset()
	poolAlibabaIcbuProductSchemaAddDraftAPIRequest.Put(v)
}
