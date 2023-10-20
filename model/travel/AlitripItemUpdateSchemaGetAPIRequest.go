package travel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripItemUpdateSchemaGetAPIRequest 获取编辑商品的schema模板 API请求
// alitrip.item.update.schema.get
//
// 获取编辑商品的schema模板。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
type AlitripItemUpdateSchemaGetAPIRequest struct {
	model.Params
	// 需要获取的编辑schema，不填默认返回全部
	_updateFieldIds []string
	// 商品id
	_itemId int64
}

// NewAlitripItemUpdateSchemaGetRequest 初始化AlitripItemUpdateSchemaGetAPIRequest对象
func NewAlitripItemUpdateSchemaGetRequest() *AlitripItemUpdateSchemaGetAPIRequest {
	return &AlitripItemUpdateSchemaGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripItemUpdateSchemaGetAPIRequest) Reset() {
	r._updateFieldIds = r._updateFieldIds[:0]
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripItemUpdateSchemaGetAPIRequest) GetApiMethodName() string {
	return "alitrip.item.update.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripItemUpdateSchemaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripItemUpdateSchemaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateFieldIds is UpdateFieldIds Setter
// 需要获取的编辑schema，不填默认返回全部
func (r *AlitripItemUpdateSchemaGetAPIRequest) SetUpdateFieldIds(_updateFieldIds []string) error {
	r._updateFieldIds = _updateFieldIds
	r.Set("update_field_ids", _updateFieldIds)
	return nil
}

// GetUpdateFieldIds UpdateFieldIds Getter
func (r AlitripItemUpdateSchemaGetAPIRequest) GetUpdateFieldIds() []string {
	return r._updateFieldIds
}

// SetItemId is ItemId Setter
// 商品id
func (r *AlitripItemUpdateSchemaGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitripItemUpdateSchemaGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolAlitripItemUpdateSchemaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripItemUpdateSchemaGetRequest()
	},
}

// GetAlitripItemUpdateSchemaGetRequest 从 sync.Pool 获取 AlitripItemUpdateSchemaGetAPIRequest
func GetAlitripItemUpdateSchemaGetAPIRequest() *AlitripItemUpdateSchemaGetAPIRequest {
	return poolAlitripItemUpdateSchemaGetAPIRequest.Get().(*AlitripItemUpdateSchemaGetAPIRequest)
}

// ReleaseAlitripItemUpdateSchemaGetAPIRequest 将 AlitripItemUpdateSchemaGetAPIRequest 放入 sync.Pool
func ReleaseAlitripItemUpdateSchemaGetAPIRequest(v *AlitripItemUpdateSchemaGetAPIRequest) {
	v.Reset()
	poolAlitripItemUpdateSchemaGetAPIRequest.Put(v)
}
