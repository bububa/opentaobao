package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitemeditschemagetAPIRequest 商品编辑获取schema信息 API请求
// alibaba.item.edit.schema.get
//
// 商品编辑时，获取商品规则信息
type AlibabaitemeditschemagetAPIRequest struct {
	model.Params
	// 制定返回schema中field字段列表，可用于裁剪返回的schema信息。不填则为全部field
	_fields []string
	// 业务扩展参数，需与平台约定好
	_bizType string
	// 商品ID
	_itemId int64
}

// NewAlibabaitemeditschemagetRequest 初始化AlibabaitemeditschemagetAPIRequest对象
func NewAlibabaitemeditschemagetRequest() *AlibabaitemeditschemagetAPIRequest {
	return &AlibabaitemeditschemagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaitemeditschemagetAPIRequest) GetApiMethodName() string {
	return "alibaba.item.edit.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaitemeditschemagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaitemeditschemagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 制定返回schema中field字段列表，可用于裁剪返回的schema信息。不填则为全部field
func (r *AlibabaitemeditschemagetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r AlibabaitemeditschemagetAPIRequest) GetFields() []string {
	return r._fields
}

// SetBizType is BizType Setter
// 业务扩展参数，需与平台约定好
func (r *AlibabaitemeditschemagetAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlibabaitemeditschemagetAPIRequest) GetBizType() string {
	return r._bizType
}

// SetItemId is ItemId Setter
// 商品ID
func (r *AlibabaitemeditschemagetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaitemeditschemagetAPIRequest) GetItemId() int64 {
	return r._itemId
}
