package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripItemSchemaUpdateAPIRequest 使用schema进行商品编辑 API请求
// alitrip.item.schema.update
//
// 飞猪度假商品使用schema进行商品编辑。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
type AlitripItemSchemaUpdateAPIRequest struct {
	model.Params
	// 商品数据
	_schemaXmlFields string
	// 商品id
	_itemId int64
}

// NewAlitripItemSchemaUpdateRequest 初始化AlitripItemSchemaUpdateAPIRequest对象
func NewAlitripItemSchemaUpdateRequest() *AlitripItemSchemaUpdateAPIRequest {
	return &AlitripItemSchemaUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripItemSchemaUpdateAPIRequest) GetApiMethodName() string {
	return "alitrip.item.schema.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripItemSchemaUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripItemSchemaUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchemaXmlFields is SchemaXmlFields Setter
// 商品数据
func (r *AlitripItemSchemaUpdateAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
	r._schemaXmlFields = _schemaXmlFields
	r.Set("schema_xml_fields", _schemaXmlFields)
	return nil
}

// GetSchemaXmlFields SchemaXmlFields Getter
func (r AlitripItemSchemaUpdateAPIRequest) GetSchemaXmlFields() string {
	return r._schemaXmlFields
}

// SetItemId is ItemId Setter
// 商品id
func (r *AlitripItemSchemaUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitripItemSchemaUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}
