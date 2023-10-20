package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemVipSchemaUpdateAPIRequest 大商家商品编辑接口 API请求
// tmall.item.vip.schema.update
//
// 大商家编辑商品
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemVipSchemaUpdateAPIRequest) Reset() {
	r._schemaXmlFields = ""
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemVipSchemaUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.item.vip.schema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemVipSchemaUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemVipSchemaUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemaXmlFields is SchemaXmlFields Setter
// 商品编辑的schema参数
func (r *TmallItemVipSchemaUpdateAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// GetSchemaXmlFields SchemaXmlFields Getter
func (r TmallItemVipSchemaUpdateAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}

// SetItemId is ItemId Setter
// 商品id
func (r *TmallItemVipSchemaUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TmallItemVipSchemaUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTmallItemVipSchemaUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemVipSchemaUpdateRequest()
	},
}

// GetTmallItemVipSchemaUpdateRequest 从 sync.Pool 获取 TmallItemVipSchemaUpdateAPIRequest
func GetTmallItemVipSchemaUpdateAPIRequest() *TmallItemVipSchemaUpdateAPIRequest {
	return poolTmallItemVipSchemaUpdateAPIRequest.Get().(*TmallItemVipSchemaUpdateAPIRequest)
}

// ReleaseTmallItemVipSchemaUpdateAPIRequest 将 TmallItemVipSchemaUpdateAPIRequest 放入 sync.Pool
func ReleaseTmallItemVipSchemaUpdateAPIRequest(v *TmallItemVipSchemaUpdateAPIRequest) {
	v.Reset()
	poolTmallItemVipSchemaUpdateAPIRequest.Put(v)
}
