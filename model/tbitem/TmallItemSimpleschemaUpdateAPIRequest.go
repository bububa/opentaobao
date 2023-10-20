package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSimpleschemaUpdateAPIRequest 天猫简化编辑商品 API请求
// tmall.item.simpleschema.update
//
// 国外大商家天猫简化编辑商品
type TmallItemSimpleschemaUpdateAPIRequest struct {
	model.Params
	// 编辑商品时提交的xml信息
	_schemaXmlFields string
	// 商品id
	_itemId int64
}

// NewTmallItemSimpleschemaUpdateRequest 初始化TmallItemSimpleschemaUpdateAPIRequest对象
func NewTmallItemSimpleschemaUpdateRequest() *TmallItemSimpleschemaUpdateAPIRequest {
	return &TmallItemSimpleschemaUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemSimpleschemaUpdateAPIRequest) Reset() {
	r._schemaXmlFields = ""
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemSimpleschemaUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.simpleschema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemSimpleschemaUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemSimpleschemaUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemaXmlFields is SchemaXmlFields Setter
// 编辑商品时提交的xml信息
func (r *TmallItemSimpleschemaUpdateAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// GetSchemaXmlFields SchemaXmlFields Getter
func (r TmallItemSimpleschemaUpdateAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallItemSimpleschemaUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemSimpleschemaUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTmallItemSimpleschemaUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemSimpleschemaUpdateRequest()
	},
}

// GetTmallItemSimpleschemaUpdateRequest 从 sync.Pool 获取 TmallItemSimpleschemaUpdateAPIRequest
func GetTmallItemSimpleschemaUpdateAPIRequest() *TmallItemSimpleschemaUpdateAPIRequest {
	return poolTmallItemSimpleschemaUpdateAPIRequest.Get().(*TmallItemSimpleschemaUpdateAPIRequest)
}

// ReleaseTmallItemSimpleschemaUpdateAPIRequest 将 TmallItemSimpleschemaUpdateAPIRequest 放入 sync.Pool
func ReleaseTmallItemSimpleschemaUpdateAPIRequest(v *TmallItemSimpleschemaUpdateAPIRequest) {
	v.Reset()
	poolTmallItemSimpleschemaUpdateAPIRequest.Put(v)
}
