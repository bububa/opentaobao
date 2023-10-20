package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemsimpleschemaupdateAPIRequest 天猫简化编辑商品 API请求
// tmall.item.simpleschema.update
//
// 国外大商家天猫简化编辑商品
type TmallitemsimpleschemaupdateAPIRequest struct {
	model.Params
	// 编辑商品时提交的xml信息
	_schemaXmlFields string
	// 商品id
	_itemId int64
}

// NewTmallitemsimpleschemaupdateRequest 初始化TmallitemsimpleschemaupdateAPIRequest对象
func NewTmallitemsimpleschemaupdateRequest() *TmallitemsimpleschemaupdateAPIRequest {
	return &TmallitemsimpleschemaupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemsimpleschemaupdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.simpleschema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemsimpleschemaupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemsimpleschemaupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemaXmlFields is SchemaXmlFields Setter
// 编辑商品时提交的xml信息
func (r *TmallitemsimpleschemaupdateAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// GetSchemaXmlFields SchemaXmlFields Getter
func (r TmallitemsimpleschemaupdateAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallitemsimpleschemaupdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallitemsimpleschemaupdateAPIRequest) GetItemId() int64 {
	return r._itemId
}
