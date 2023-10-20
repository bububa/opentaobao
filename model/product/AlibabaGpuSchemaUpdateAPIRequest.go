package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabagpuschemaupdateAPIRequest 产品更新接口 API请求
// alibaba.gpu.schema.update
//
// 产品更新接口
type AlibabagpuschemaupdateAPIRequest struct {
	model.Params
	// 更新产品提交的schema数据
	_schemaXmlFields string
	// 产品ID
	_productId int64
	// 当前用户所在渠道如0代表天猫，8代表淘宝
	_providerId int64
}

// NewAlibabagpuschemaupdateRequest 初始化AlibabagpuschemaupdateAPIRequest对象
func NewAlibabagpuschemaupdateRequest() *AlibabagpuschemaupdateAPIRequest {
	return &AlibabagpuschemaupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabagpuschemaupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.gpu.schema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabagpuschemaupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabagpuschemaupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemaXmlFields is SchemaXmlFields Setter
// 更新产品提交的schema数据
func (r *AlibabagpuschemaupdateAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// GetSchemaXmlFields SchemaXmlFields Getter
func (r AlibabagpuschemaupdateAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}

// SetProductId is ProductId Setter
// 产品ID
func (r *AlibabagpuschemaupdateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabagpuschemaupdateAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetProviderId is ProviderId Setter
// 当前用户所在渠道如0代表天猫，8代表淘宝
func (r *AlibabagpuschemaupdateAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlibabagpuschemaupdateAPIRequest) GetProviderId() int64 {
	return r._providerId
}
