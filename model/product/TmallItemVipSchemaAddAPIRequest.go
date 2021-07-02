package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallItemVipSchemaAddAPIRequest 大商家商品发布接口 API请求
// tmall.item.vip.schema.add
//
// 大商家商品发布接口
type TmallItemVipSchemaAddAPIRequest struct {
	model.Params
	// 商品发布schema参数
	_schemaXmlFields string
}

// NewTmallItemVipSchemaAddRequest 初始化TmallItemVipSchemaAddAPIRequest对象
func NewTmallItemVipSchemaAddRequest() *TmallItemVipSchemaAddAPIRequest {
	return &TmallItemVipSchemaAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemVipSchemaAddAPIRequest) GetApiMethodName() string {
	return "tmall.item.vip.schema.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemVipSchemaAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSchemaXmlFields is SchemaXmlFields Setter
// 商品发布schema参数
func (r *TmallItemVipSchemaAddAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// GetSchemaXmlFields SchemaXmlFields Getter
func (r TmallItemVipSchemaAddAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}
