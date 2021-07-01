package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenItemsMarkingAPIRequest
商品通自动打标 API请求
taobao.qimen.items.marking

调用该接口，对商品进行XXXX标的打标、去标的动作。 */
type TaobaoQimenItemsMarkingAPIRequest struct {
	model.Params
	// 操作类型，string（50），ADD=打标，DELETE=去标，必填
	_actionType string
	// 打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填
	_tagType string
	// 线上商品ID，long，必填
	_itemIds []int64
	// 备注，string（500）
	_remark string
}

// NewTaobaoQimenItemsMarkingRequest 初始化TaobaoQimenItemsMarkingAPIRequest对象
func NewTaobaoQimenItemsMarkingRequest() *TaobaoQimenItemsMarkingAPIRequest {
	return &TaobaoQimenItemsMarkingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemsMarkingAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.items.marking"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemsMarkingAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActionType Setter
// 操作类型，string（50），ADD=打标，DELETE=去标，必填
func (r *TaobaoQimenItemsMarkingAPIRequest) SetActionType(_actionType string) error {
	r._actionType = _actionType
	r.Set("action_type", _actionType)
	return nil
}

// Get ActionType Getter
func (r TaobaoQimenItemsMarkingAPIRequest) GetActionType() string {
	return r._actionType
}

// Set is TagType Setter
// 打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填
func (r *TaobaoQimenItemsMarkingAPIRequest) SetTagType(_tagType string) error {
	r._tagType = _tagType
	r.Set("tag_type", _tagType)
	return nil
}

// Get TagType Getter
func (r TaobaoQimenItemsMarkingAPIRequest) GetTagType() string {
	return r._tagType
}

// Set is ItemIds Setter
// 线上商品ID，long，必填
func (r *TaobaoQimenItemsMarkingAPIRequest) SetItemIds(_itemIds []int64) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// Get ItemIds Getter
func (r TaobaoQimenItemsMarkingAPIRequest) GetItemIds() []int64 {
	return r._itemIds
}

// Set is Remark Setter
// 备注，string（500）
func (r *TaobaoQimenItemsMarkingAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// Get Remark Getter
func (r TaobaoQimenItemsMarkingAPIRequest) GetRemark() string {
	return r._remark
}
