package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvitemrechargeeditAPIRequest 闲鱼商品直充功能编辑 API请求
// alibaba.idle.isv.item.recharge.edit
//
// 闲鱼商品直充功能编辑
type AlibabaidleisvitemrechargeeditAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 直充信息
	_idleItemApiAutoRechargeDo *IdleItemApiAutoRechargeDo
}

// NewAlibabaidleisvitemrechargeeditRequest 初始化AlibabaidleisvitemrechargeeditAPIRequest对象
func NewAlibabaidleisvitemrechargeeditRequest() *AlibabaidleisvitemrechargeeditAPIRequest {
	return &AlibabaidleisvitemrechargeeditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleisvitemrechargeeditAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.item.recharge.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleisvitemrechargeeditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleisvitemrechargeeditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *AlibabaidleisvitemrechargeeditAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaidleisvitemrechargeeditAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetIdleItemApiAutoRechargeDo is IdleItemApiAutoRechargeDo Setter
// 直充信息
func (r *AlibabaidleisvitemrechargeeditAPIRequest) SetIdleItemApiAutoRechargeDo(_idleItemApiAutoRechargeDo *IdleItemApiAutoRechargeDo) error {
	r._idleItemApiAutoRechargeDo = _idleItemApiAutoRechargeDo
	r.Set("idle_item_api_auto_recharge_do", _idleItemApiAutoRechargeDo)
	return nil
}

// GetIdleItemApiAutoRechargeDo IdleItemApiAutoRechargeDo Getter
func (r AlibabaidleisvitemrechargeeditAPIRequest) GetIdleItemApiAutoRechargeDo() *IdleItemApiAutoRechargeDo {
	return r._idleItemApiAutoRechargeDo
}
