package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopentradetoolsitemsbindAPIRequest 交易开放商品绑定 API请求
// taobao.opentrade.tools.items.bind
//
// 交易开放商品绑定
type TaobaoopentradetoolsitemsbindAPIRequest struct {
	model.Params
	// 待绑定商品id
	_itemIds []int64
	// 绑定交易开放场景的C端小程序ID
	_miniappId int64
}

// NewTaobaoopentradetoolsitemsbindRequest 初始化TaobaoopentradetoolsitemsbindAPIRequest对象
func NewTaobaoopentradetoolsitemsbindRequest() *TaobaoopentradetoolsitemsbindAPIRequest {
	return &TaobaoopentradetoolsitemsbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopentradetoolsitemsbindAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.tools.items.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopentradetoolsitemsbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopentradetoolsitemsbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 待绑定商品id
func (r *TaobaoopentradetoolsitemsbindAPIRequest) SetItemIds(_itemIds []int64) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoopentradetoolsitemsbindAPIRequest) GetItemIds() []int64 {
	return r._itemIds
}

// SetMiniappId is MiniappId Setter
// 绑定交易开放场景的C端小程序ID
func (r *TaobaoopentradetoolsitemsbindAPIRequest) SetMiniappId(_miniappId int64) error {
	r._miniappId = _miniappId
	r.Set("miniapp_id", _miniappId)
	return nil
}

// GetMiniappId MiniappId Getter
func (r TaobaoopentradetoolsitemsbindAPIRequest) GetMiniappId() int64 {
	return r._miniappId
}
