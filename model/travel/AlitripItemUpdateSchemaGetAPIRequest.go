package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripitemupdateschemagetAPIRequest 获取编辑商品的schema模板 API请求
// alitrip.item.update.schema.get
//
// 获取编辑商品的schema模板。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
type AlitripitemupdateschemagetAPIRequest struct {
	model.Params
	// 需要获取的编辑schema，不填默认返回全部
	_updateFieldIds []string
	// 商品id
	_itemId int64
}

// NewAlitripitemupdateschemagetRequest 初始化AlitripitemupdateschemagetAPIRequest对象
func NewAlitripitemupdateschemagetRequest() *AlitripitemupdateschemagetAPIRequest {
	return &AlitripitemupdateschemagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripitemupdateschemagetAPIRequest) GetApiMethodName() string {
	return "alitrip.item.update.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripitemupdateschemagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripitemupdateschemagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateFieldIds is UpdateFieldIds Setter
// 需要获取的编辑schema，不填默认返回全部
func (r *AlitripitemupdateschemagetAPIRequest) SetUpdateFieldIds(_updateFieldIds []string) error {
	r._updateFieldIds = _updateFieldIds
	r.Set("update_field_ids", _updateFieldIds)
	return nil
}

// GetUpdateFieldIds UpdateFieldIds Getter
func (r AlitripitemupdateschemagetAPIRequest) GetUpdateFieldIds() []string {
	return r._updateFieldIds
}

// SetItemId is ItemId Setter
// 商品id
func (r *AlitripitemupdateschemagetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitripitemupdateschemagetAPIRequest) GetItemId() int64 {
	return r._itemId
}
