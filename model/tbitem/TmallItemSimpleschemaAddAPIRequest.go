package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemsimpleschemaaddAPIRequest 天猫简化发布商品 API请求
// tmall.item.simpleschema.add
//
// 天猫简化版schema发布商品。
type TmallitemsimpleschemaaddAPIRequest struct {
	model.Params
	// 根据tmall.item.add.simpleschema.get生成的商品发布规则入参数据
	_schemaXmlFields string
}

// NewTmallitemsimpleschemaaddRequest 初始化TmallitemsimpleschemaaddAPIRequest对象
func NewTmallitemsimpleschemaaddRequest() *TmallitemsimpleschemaaddAPIRequest {
	return &TmallitemsimpleschemaaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemsimpleschemaaddAPIRequest) GetApiMethodName() string {
	return "tmall.item.simpleschema.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemsimpleschemaaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemsimpleschemaaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemaXmlFields is SchemaXmlFields Setter
// 根据tmall.item.add.simpleschema.get生成的商品发布规则入参数据
func (r *TmallitemsimpleschemaaddAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// GetSchemaXmlFields SchemaXmlFields Getter
func (r TmallitemsimpleschemaaddAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}
