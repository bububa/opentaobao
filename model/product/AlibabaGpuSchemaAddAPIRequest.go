package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabagpuschemaaddAPIRequest 使用schema文件发布产品 API请求
// alibaba.gpu.schema.add
//
// 使用Schema文件发布一个产品
type AlibabagpuschemaaddAPIRequest struct {
	model.Params
	// 根据alibaba.gpu.add.schema.get获取的规则提交上来的schema
	_schemaXmlFields string
	// 叶子类目ID
	_leafCatId int64
	// 品牌ID
	_brandId int64
	// 当前用户所在渠道如0代表天猫，8代表淘宝
	_providerId int64
}

// NewAlibabagpuschemaaddRequest 初始化AlibabagpuschemaaddAPIRequest对象
func NewAlibabagpuschemaaddRequest() *AlibabagpuschemaaddAPIRequest {
	return &AlibabagpuschemaaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabagpuschemaaddAPIRequest) GetApiMethodName() string {
	return "alibaba.gpu.schema.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabagpuschemaaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabagpuschemaaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemaXmlFields is SchemaXmlFields Setter
// 根据alibaba.gpu.add.schema.get获取的规则提交上来的schema
func (r *AlibabagpuschemaaddAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// GetSchemaXmlFields SchemaXmlFields Getter
func (r AlibabagpuschemaaddAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}

// SetLeafCatId is LeafCatId Setter
// 叶子类目ID
func (r *AlibabagpuschemaaddAPIRequest) SetLeafCatId(_leafCatId int64) error {
	r._leafCatId = _leafCatId
	r.Set("leaf_cat_id", _leafCatId)
	return nil
}

// GetLeafCatId LeafCatId Getter
func (r AlibabagpuschemaaddAPIRequest) GetLeafCatId() int64 {
	return r._leafCatId
}

// SetBrandId is BrandId Setter
// 品牌ID
func (r *AlibabagpuschemaaddAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r AlibabagpuschemaaddAPIRequest) GetBrandId() int64 {
	return r._brandId
}

// SetProviderId is ProviderId Setter
// 当前用户所在渠道如0代表天猫，8代表淘宝
func (r *AlibabagpuschemaaddAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlibabagpuschemaaddAPIRequest) GetProviderId() int64 {
	return r._providerId
}
