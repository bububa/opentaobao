package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemSimpleschemaAddAPIRequest
天猫简化发布商品 API请求
tmall.item.simpleschema.add

天猫简化版schema发布商品。 */
type TmallItemSimpleschemaAddAPIRequest struct {
	model.Params
	// 根据tmall.item.add.simpleschema.get生成的商品发布规则入参数据
	_schemaXmlFields string
}

// NewTmallItemSimpleschemaAddRequest 初始化TmallItemSimpleschemaAddAPIRequest对象
func NewTmallItemSimpleschemaAddRequest() *TmallItemSimpleschemaAddAPIRequest {
	return &TmallItemSimpleschemaAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSimpleschemaAddAPIRequest) GetApiMethodName() string {
	return "tmall.item.simpleschema.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSimpleschemaAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SchemaXmlFields Setter
// 根据tmall.item.add.simpleschema.get生成的商品发布规则入参数据
func (r *TmallItemSimpleschemaAddAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// Get SchemaXmlFields Getter
func (r TmallItemSimpleschemaAddAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}
