package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopentradetoolsitemsunbindAPIRequest 交易开放商品解绑 API请求
// taobao.opentrade.tools.items.unbind
//
// 交易开放商品解绑
type TaobaoopentradetoolsitemsunbindAPIRequest struct {
	model.Params
	// 商品id
	_itemIds []int64
	// 绑定交易开放场景的C端小程序ID
	_miniappId int64
}

// NewTaobaoopentradetoolsitemsunbindRequest 初始化TaobaoopentradetoolsitemsunbindAPIRequest对象
func NewTaobaoopentradetoolsitemsunbindRequest() *TaobaoopentradetoolsitemsunbindAPIRequest {
	return &TaobaoopentradetoolsitemsunbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopentradetoolsitemsunbindAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.tools.items.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopentradetoolsitemsunbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopentradetoolsitemsunbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 商品id
func (r *TaobaoopentradetoolsitemsunbindAPIRequest) SetItemIds(_itemIds []int64) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoopentradetoolsitemsunbindAPIRequest) GetItemIds() []int64 {
	return r._itemIds
}

// SetMiniappId is MiniappId Setter
// 绑定交易开放场景的C端小程序ID
func (r *TaobaoopentradetoolsitemsunbindAPIRequest) SetMiniappId(_miniappId int64) error {
	r._miniappId = _miniappId
	r.Set("miniapp_id", _miniappId)
	return nil
}

// GetMiniappId MiniappId Getter
func (r TaobaoopentradetoolsitemsunbindAPIRequest) GetMiniappId() int64 {
	return r._miniappId
}
