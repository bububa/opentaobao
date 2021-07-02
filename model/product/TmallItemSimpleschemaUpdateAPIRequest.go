package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSimpleschemaUpdateAPIRequest 天猫简化编辑商品 API请求
// tmall.item.simpleschema.update
//
// 国外大商家天猫简化编辑商品
type TmallItemSimpleschemaUpdateAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 编辑商品时提交的xml信息
	_schemaXmlFields string
}

// NewTmallItemSimpleschemaUpdateRequest 初始化TmallItemSimpleschemaUpdateAPIRequest对象
func NewTmallItemSimpleschemaUpdateRequest() *TmallItemSimpleschemaUpdateAPIRequest {
	return &TmallItemSimpleschemaUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSimpleschemaUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.simpleschema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSimpleschemaUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemId Setter
// 商品id
func (r *TmallItemSimpleschemaUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TmallItemSimpleschemaUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is SchemaXmlFields Setter
// 编辑商品时提交的xml信息
func (r *TmallItemSimpleschemaUpdateAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// Get SchemaXmlFields Getter
func (r TmallItemSimpleschemaUpdateAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}
