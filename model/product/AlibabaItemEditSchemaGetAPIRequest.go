package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItemEditSchemaGetAPIRequest 商品编辑获取schema信息 API请求
// alibaba.item.edit.schema.get
//
// 商品编辑时，获取商品规则信息
type AlibabaItemEditSchemaGetAPIRequest struct {
	model.Params
	// 业务扩展参数，需与平台约定好
	_bizType string
	// 商品ID
	_itemId int64
	// 制定返回schema中field字段列表，可用于裁剪返回的schema信息。不填则为全部field
	_fields []string
}

// NewAlibabaItemEditSchemaGetRequest 初始化AlibabaItemEditSchemaGetAPIRequest对象
func NewAlibabaItemEditSchemaGetRequest() *AlibabaItemEditSchemaGetAPIRequest {
	return &AlibabaItemEditSchemaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItemEditSchemaGetAPIRequest) GetApiMethodName() string {
	return "alibaba.item.edit.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItemEditSchemaGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBizType is BizType Setter
// 业务扩展参数，需与平台约定好
func (r *AlibabaItemEditSchemaGetAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabaItemEditSchemaGetAPIRequest) GetBizType() string {
	return r._bizType
}

// SetItemId is ItemId Setter
// 商品ID
func (r *AlibabaItemEditSchemaGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaItemEditSchemaGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetFields is Fields Setter
// 制定返回schema中field字段列表，可用于裁剪返回的schema信息。不填则为全部field
func (r *AlibabaItemEditSchemaGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r AlibabaItemEditSchemaGetAPIRequest) GetFields() []string {
	return r._fields
}
