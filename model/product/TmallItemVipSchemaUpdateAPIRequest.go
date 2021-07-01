package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemVipSchemaUpdateAPIRequest
大商家商品编辑接口 API请求
tmall.item.vip.schema.update

大商家编辑商品 */
type TmallItemVipSchemaUpdateAPIRequest struct {
	model.Params
	// 商品编辑的schema参数
	_schemaXmlFields string
	// 商品id
	_itemId int64
}

// NewTmallItemVipSchemaUpdateRequest 初始化TmallItemVipSchemaUpdateAPIRequest对象
func NewTmallItemVipSchemaUpdateRequest() *TmallItemVipSchemaUpdateAPIRequest {
	return &TmallItemVipSchemaUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemVipSchemaUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.vip.schema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemVipSchemaUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SchemaXmlFields Setter
// 商品编辑的schema参数
func (r *TmallItemVipSchemaUpdateAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// Get SchemaXmlFields Getter
func (r TmallItemVipSchemaUpdateAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}

// Set is ItemId Setter
// 商品id
func (r *TmallItemVipSchemaUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TmallItemVipSchemaUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}
