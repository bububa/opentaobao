package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaGpuSchemaAddAPIRequest
使用schema文件发布产品 API请求
alibaba.gpu.schema.add

使用Schema文件发布一个产品 */
type AlibabaGpuSchemaAddAPIRequest struct {
	model.Params
	// 叶子类目ID
	_leafCatId int64
	// 品牌ID
	_brandId int64
	// 根据alibaba.gpu.add.schema.get获取的规则提交上来的schema
	_schemaXmlFields string
	// 当前用户所在渠道如0代表天猫，8代表淘宝
	_providerId int64
}

// NewAlibabaGpuSchemaAddRequest 初始化AlibabaGpuSchemaAddAPIRequest对象
func NewAlibabaGpuSchemaAddRequest() *AlibabaGpuSchemaAddAPIRequest {
	return &AlibabaGpuSchemaAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaGpuSchemaAddAPIRequest) GetApiMethodName() string {
	return "alibaba.gpu.schema.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaGpuSchemaAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LeafCatId Setter
// 叶子类目ID
func (r *AlibabaGpuSchemaAddAPIRequest) SetLeafCatId(_leafCatId int64) error {
	r._leafCatId = _leafCatId
	r.Set("leaf_cat_id", _leafCatId)
	return nil
}

// Get LeafCatId Getter
func (r AlibabaGpuSchemaAddAPIRequest) GetLeafCatId() int64 {
	return r._leafCatId
}

// Set is BrandId Setter
// 品牌ID
func (r *AlibabaGpuSchemaAddAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// Get BrandId Getter
func (r AlibabaGpuSchemaAddAPIRequest) GetBrandId() int64 {
	return r._brandId
}

// Set is SchemaXmlFields Setter
// 根据alibaba.gpu.add.schema.get获取的规则提交上来的schema
func (r *AlibabaGpuSchemaAddAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// Get SchemaXmlFields Getter
func (r AlibabaGpuSchemaAddAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}

// Set is ProviderId Setter
// 当前用户所在渠道如0代表天猫，8代表淘宝
func (r *AlibabaGpuSchemaAddAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// Get ProviderId Getter
func (r AlibabaGpuSchemaAddAPIRequest) GetProviderId() int64 {
	return r._providerId
}
