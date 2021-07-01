package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripItemUpdateSchemaGetAPIRequest
获取编辑商品的schema模板 API请求
alitrip.item.update.schema.get

获取编辑商品的schema模板。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003) */
type AlitripItemUpdateSchemaGetAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 需要获取的编辑schema，不填默认返回全部
	_updateFieldIds []string
}

// NewAlitripItemUpdateSchemaGetRequest 初始化AlitripItemUpdateSchemaGetAPIRequest对象
func NewAlitripItemUpdateSchemaGetRequest() *AlitripItemUpdateSchemaGetAPIRequest {
	return &AlitripItemUpdateSchemaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripItemUpdateSchemaGetAPIRequest) GetApiMethodName() string {
	return "alitrip.item.update.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripItemUpdateSchemaGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemId Setter
// 商品id
func (r *AlitripItemUpdateSchemaGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r AlitripItemUpdateSchemaGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is UpdateFieldIds Setter
// 需要获取的编辑schema，不填默认返回全部
func (r *AlitripItemUpdateSchemaGetAPIRequest) SetUpdateFieldIds(_updateFieldIds []string) error {
	r._updateFieldIds = _updateFieldIds
	r.Set("update_field_ids", _updateFieldIds)
	return nil
}

// Get UpdateFieldIds Getter
func (r AlitripItemUpdateSchemaGetAPIRequest) GetUpdateFieldIds() []string {
	return r._updateFieldIds
}
