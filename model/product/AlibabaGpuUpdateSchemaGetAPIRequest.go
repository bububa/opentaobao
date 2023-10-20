package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaGpuUpdateSchemaGetAPIRequest 获取产品编辑schema规则的接口 API请求
// alibaba.gpu.update.schema.get
//
// 获取产品编辑schema规则的接口
type AlibabaGpuUpdateSchemaGetAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
	// 当前用户所在渠道如0代表天猫，8代表淘宝
	_providerId int64
}

// NewAlibabaGpuUpdateSchemaGetRequest 初始化AlibabaGpuUpdateSchemaGetAPIRequest对象
func NewAlibabaGpuUpdateSchemaGetRequest() *AlibabaGpuUpdateSchemaGetAPIRequest {
	return &AlibabaGpuUpdateSchemaGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaGpuUpdateSchemaGetAPIRequest) Reset() {
	r._productId = 0
	r._providerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaGpuUpdateSchemaGetAPIRequest) GetApiMethodName() string {
	return "alibaba.gpu.update.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaGpuUpdateSchemaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaGpuUpdateSchemaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品ID
func (r *AlibabaGpuUpdateSchemaGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabaGpuUpdateSchemaGetAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetProviderId is ProviderId Setter
// 当前用户所在渠道如0代表天猫，8代表淘宝
func (r *AlibabaGpuUpdateSchemaGetAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlibabaGpuUpdateSchemaGetAPIRequest) GetProviderId() int64 {
	return r._providerId
}

var poolAlibabaGpuUpdateSchemaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaGpuUpdateSchemaGetRequest()
	},
}

// GetAlibabaGpuUpdateSchemaGetRequest 从 sync.Pool 获取 AlibabaGpuUpdateSchemaGetAPIRequest
func GetAlibabaGpuUpdateSchemaGetAPIRequest() *AlibabaGpuUpdateSchemaGetAPIRequest {
	return poolAlibabaGpuUpdateSchemaGetAPIRequest.Get().(*AlibabaGpuUpdateSchemaGetAPIRequest)
}

// ReleaseAlibabaGpuUpdateSchemaGetAPIRequest 将 AlibabaGpuUpdateSchemaGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaGpuUpdateSchemaGetAPIRequest(v *AlibabaGpuUpdateSchemaGetAPIRequest) {
	v.Reset()
	poolAlibabaGpuUpdateSchemaGetAPIRequest.Put(v)
}
