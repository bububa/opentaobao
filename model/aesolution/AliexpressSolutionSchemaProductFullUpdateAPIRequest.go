package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionschemaproductfullupdateAPIRequest aliexpress.solution.schema.product.full.update API请求
// aliexpress.solution.schema.product.full.update
//
// Schema interface for product full update. QPS(Invoke per second) for this API is limited to 100 for each appkey and 50 for each seller.
type AliexpresssolutionschemaproductfullupdateAPIRequest struct {
	model.Params
	// Product full update request. To learn how to generate the content, please refer to https://developers.aliexpress.com/en/doc.htm?docId=109760&docType=1.  Be aware that the aliexpress_product_id field should be replaced by the product ID belonged to the seller.
	_schemaFullUpdateRequest string
	// More information of the request.
	_developerFeatures string
}

// NewAliexpresssolutionschemaproductfullupdateRequest 初始化AliexpresssolutionschemaproductfullupdateAPIRequest对象
func NewAliexpresssolutionschemaproductfullupdateRequest() *AliexpresssolutionschemaproductfullupdateAPIRequest {
	return &AliexpresssolutionschemaproductfullupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionschemaproductfullupdateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.schema.product.full.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionschemaproductfullupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionschemaproductfullupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemaFullUpdateRequest is SchemaFullUpdateRequest Setter
// Product full update request. To learn how to generate the content, please refer to https://developers.aliexpress.com/en/doc.htm?docId=109760&amp;docType=1.  Be aware that the aliexpress_product_id field should be replaced by the product ID belonged to the seller.
func (r *AliexpresssolutionschemaproductfullupdateAPIRequest) SetSchemaFullUpdateRequest(_schemaFullUpdateRequest string) error {
	r._schemaFullUpdateRequest = _schemaFullUpdateRequest
	r.Set("schema_full_update_request", _schemaFullUpdateRequest)
	return nil
}

// GetSchemaFullUpdateRequest SchemaFullUpdateRequest Getter
func (r AliexpresssolutionschemaproductfullupdateAPIRequest) GetSchemaFullUpdateRequest() string {
	return r._schemaFullUpdateRequest
}

// SetDeveloperFeatures is DeveloperFeatures Setter
// More information of the request.
func (r *AliexpresssolutionschemaproductfullupdateAPIRequest) SetDeveloperFeatures(_developerFeatures string) error {
	r._developerFeatures = _developerFeatures
	r.Set("developer_features", _developerFeatures)
	return nil
}

// GetDeveloperFeatures DeveloperFeatures Getter
func (r AliexpresssolutionschemaproductfullupdateAPIRequest) GetDeveloperFeatures() string {
	return r._developerFeatures
}
