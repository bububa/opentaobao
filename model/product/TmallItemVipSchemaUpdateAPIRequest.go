package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemvipschemaupdateAPIRequest 大商家商品编辑接口 API请求
// tmall.item.vip.schema.update
//
// 大商家编辑商品
type TmallitemvipschemaupdateAPIRequest struct {
	model.Params
	// 商品编辑的schema参数
	_schemaXmlFields string
	// 商品id
	_itemId int64
}

// NewTmallitemvipschemaupdateRequest 初始化TmallitemvipschemaupdateAPIRequest对象
func NewTmallitemvipschemaupdateRequest() *TmallitemvipschemaupdateAPIRequest {
	return &TmallitemvipschemaupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemvipschemaupdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.vip.schema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemvipschemaupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemvipschemaupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemaXmlFields is SchemaXmlFields Setter
// 商品编辑的schema参数
func (r *TmallitemvipschemaupdateAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// GetSchemaXmlFields SchemaXmlFields Getter
func (r TmallitemvipschemaupdateAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallitemvipschemaupdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemvipschemaupdateAPIRequest) GetItemId() int64 {
	return r._itemId
}
