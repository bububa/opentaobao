package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaGpuSchemaUpdateAPIRequest 产品更新接口 API请求
// alibaba.gpu.schema.update
//
// 产品更新接口
type AlibabaGpuSchemaUpdateAPIRequest struct {
	model.Params
	// 更新产品提交的schema数据
	_schemaXmlFields string
	// 产品ID
	_productId int64
	// 当前用户所在渠道如0代表天猫，8代表淘宝
	_providerId int64
}

// NewAlibabaGpuSchemaUpdateRequest 初始化AlibabaGpuSchemaUpdateAPIRequest对象
func NewAlibabaGpuSchemaUpdateRequest() *AlibabaGpuSchemaUpdateAPIRequest {
	return &AlibabaGpuSchemaUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaGpuSchemaUpdateAPIRequest) Reset() {
	r._schemaXmlFields = ""
	r._productId = 0
	r._providerId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaGpuSchemaUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.gpu.schema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaGpuSchemaUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaGpuSchemaUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemaXmlFields is SchemaXmlFields Setter
// 更新产品提交的schema数据
func (r *AlibabaGpuSchemaUpdateAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// GetSchemaXmlFields SchemaXmlFields Getter
func (r AlibabaGpuSchemaUpdateAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}

// SetProductId is ProductId Setter
// 产品ID
func (r *AlibabaGpuSchemaUpdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabaGpuSchemaUpdateAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetProviderId is ProviderId Setter
// 当前用户所在渠道如0代表天猫，8代表淘宝
func (r *AlibabaGpuSchemaUpdateAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlibabaGpuSchemaUpdateAPIRequest) GetProviderId() int64 {
	return r._providerId
}

var poolAlibabaGpuSchemaUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaGpuSchemaUpdateRequest()
	},
}

// GetAlibabaGpuSchemaUpdateRequest 从 sync.Pool 获取 AlibabaGpuSchemaUpdateAPIRequest
func GetAlibabaGpuSchemaUpdateAPIRequest() *AlibabaGpuSchemaUpdateAPIRequest {
	return poolAlibabaGpuSchemaUpdateAPIRequest.Get().(*AlibabaGpuSchemaUpdateAPIRequest)
}

// ReleaseAlibabaGpuSchemaUpdateAPIRequest 将 AlibabaGpuSchemaUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaGpuSchemaUpdateAPIRequest(v *AlibabaGpuSchemaUpdateAPIRequest) {
	v.Reset()
	poolAlibabaGpuSchemaUpdateAPIRequest.Put(v)
}
