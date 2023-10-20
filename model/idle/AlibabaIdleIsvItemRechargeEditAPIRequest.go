package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvItemRechargeEditAPIRequest 闲鱼商品直充功能编辑 API请求
// alibaba.idle.isv.item.recharge.edit
//
// 闲鱼商品直充功能编辑
type AlibabaIdleIsvItemRechargeEditAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 直充信息
	_idleItemApiAutoRechargeDo *IdleItemApiAutoRechargeDo
}

// NewAlibabaIdleIsvItemRechargeEditRequest 初始化AlibabaIdleIsvItemRechargeEditAPIRequest对象
func NewAlibabaIdleIsvItemRechargeEditRequest() *AlibabaIdleIsvItemRechargeEditAPIRequest {
	return &AlibabaIdleIsvItemRechargeEditAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleIsvItemRechargeEditAPIRequest) Reset() {
	r._itemId = 0
	r._idleItemApiAutoRechargeDo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvItemRechargeEditAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.item.recharge.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvItemRechargeEditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleIsvItemRechargeEditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 商品id
func (r *AlibabaIdleIsvItemRechargeEditAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaIdleIsvItemRechargeEditAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetIdleItemApiAutoRechargeDo is IdleItemApiAutoRechargeDo Setter
// 直充信息
func (r *AlibabaIdleIsvItemRechargeEditAPIRequest) SetIdleItemApiAutoRechargeDo(_idleItemApiAutoRechargeDo *IdleItemApiAutoRechargeDo) error {
	r._idleItemApiAutoRechargeDo = _idleItemApiAutoRechargeDo
	r.Set("idle_item_api_auto_recharge_do", _idleItemApiAutoRechargeDo)
	return nil
}

// GetIdleItemApiAutoRechargeDo IdleItemApiAutoRechargeDo Getter
func (r AlibabaIdleIsvItemRechargeEditAPIRequest) GetIdleItemApiAutoRechargeDo() *IdleItemApiAutoRechargeDo {
	return r._idleItemApiAutoRechargeDo
}

var poolAlibabaIdleIsvItemRechargeEditAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleIsvItemRechargeEditRequest()
	},
}

// GetAlibabaIdleIsvItemRechargeEditRequest 从 sync.Pool 获取 AlibabaIdleIsvItemRechargeEditAPIRequest
func GetAlibabaIdleIsvItemRechargeEditAPIRequest() *AlibabaIdleIsvItemRechargeEditAPIRequest {
	return poolAlibabaIdleIsvItemRechargeEditAPIRequest.Get().(*AlibabaIdleIsvItemRechargeEditAPIRequest)
}

// ReleaseAlibabaIdleIsvItemRechargeEditAPIRequest 将 AlibabaIdleIsvItemRechargeEditAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleIsvItemRechargeEditAPIRequest(v *AlibabaIdleIsvItemRechargeEditAPIRequest) {
	v.Reset()
	poolAlibabaIdleIsvItemRechargeEditAPIRequest.Put(v)
}
