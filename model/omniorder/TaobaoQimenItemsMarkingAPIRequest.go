package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenitemsmarkingAPIRequest 商品通自动打标 API请求
// taobao.qimen.items.marking
//
// 调用该接口，对商品进行XXXX标的打标、去标的动作。
type TaobaoqimenitemsmarkingAPIRequest struct {
	model.Params
	// 线上商品ID，long，必填
	_itemIds []string
	// 操作类型，string（50），ADD=打标，DELETE=去标，必填
	_actionType string
	// 打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填
	_tagType string
	// 备注，string（500）
	_remark string
}

// NewTaobaoqimenitemsmarkingRequest 初始化TaobaoqimenitemsmarkingAPIRequest对象
func NewTaobaoqimenitemsmarkingRequest() *TaobaoqimenitemsmarkingAPIRequest {
	return &TaobaoqimenitemsmarkingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenitemsmarkingAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.items.marking"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenitemsmarkingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenitemsmarkingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 线上商品ID，long，必填
func (r *TaobaoqimenitemsmarkingAPIRequest) SetItemIds(_itemIds []string) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoqimenitemsmarkingAPIRequest) GetItemIds() []string {
	return r._itemIds
}

// SetActionType is ActionType Setter
// 操作类型，string（50），ADD=打标，DELETE=去标，必填
func (r *TaobaoqimenitemsmarkingAPIRequest) SetActionType(_actionType string) error {
	r._actionType = _actionType
	r.Set("action_type", _actionType)
	return nil
}

// GetActionType ActionType Getter
func (r TaobaoqimenitemsmarkingAPIRequest) GetActionType() string {
	return r._actionType
}

// SetTagType is TagType Setter
// 打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填
func (r *TaobaoqimenitemsmarkingAPIRequest) SetTagType(_tagType string) error {
	r._tagType = _tagType
	r.Set("tag_type", _tagType)
	return nil
}

// GetTagType TagType Getter
func (r TaobaoqimenitemsmarkingAPIRequest) GetTagType() string {
	return r._tagType
}

// SetRemark is Remark Setter
// 备注，string（500）
func (r *TaobaoqimenitemsmarkingAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoqimenitemsmarkingAPIRequest) GetRemark() string {
	return r._remark
}
