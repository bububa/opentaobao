package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemvipschemaaddAPIRequest 大商家商品发布接口 API请求
// tmall.item.vip.schema.add
//
// 大商家商品发布接口
type TmallitemvipschemaaddAPIRequest struct {
	model.Params
	// 商品发布schema参数
	_schemaXmlFields string
}

// NewTmallitemvipschemaaddRequest 初始化TmallitemvipschemaaddAPIRequest对象
func NewTmallitemvipschemaaddRequest() *TmallitemvipschemaaddAPIRequest {
	return &TmallitemvipschemaaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemvipschemaaddAPIRequest) GetApiMethodName() string {
	return "tmall.item.vip.schema.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemvipschemaaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemvipschemaaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemaXmlFields is SchemaXmlFields Setter
// 商品发布schema参数
func (r *TmallitemvipschemaaddAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// GetSchemaXmlFields SchemaXmlFields Getter
func (r TmallitemvipschemaaddAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}
