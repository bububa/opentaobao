package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemCarturlGetAPIRequest 加购URL获取 API请求
// taobao.item.carturl.get
//
// 获取加购URL，支持添加商品到购物车
type TaobaoItemCarturlGetAPIRequest struct {
	model.Params
	// 商品信息，格式为 商品ID_SKU ID_数量，多条记录以逗号(,)分割
	_itemIds []string
	// 回调地址，需要是EWS域名地址。可不填，默认到购物车页面
	_callbackUrl string
	// 商家Nick，优先使用user_id
	_userNick string
	// 扩展属性，关注店铺的时候会传递下去，格式为K:V|K:V格式
	_extParams string
	// 端类型，默认是tb，可选tb,hm
	_type string
	// 商家ID
	_userId int64
}

// NewTaobaoItemCarturlGetRequest 初始化TaobaoItemCarturlGetAPIRequest对象
func NewTaobaoItemCarturlGetRequest() *TaobaoItemCarturlGetAPIRequest {
	return &TaobaoItemCarturlGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemCarturlGetAPIRequest) GetApiMethodName() string {
	return "taobao.item.carturl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemCarturlGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemCarturlGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 商品信息，格式为 商品ID_SKU ID_数量，多条记录以逗号(,)分割
func (r *TaobaoItemCarturlGetAPIRequest) SetItemIds(_itemIds []string) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TaobaoItemCarturlGetAPIRequest) GetItemIds() []string {
	return r._itemIds
}

// SetCallbackUrl is CallbackUrl Setter
// 回调地址，需要是EWS域名地址。可不填，默认到购物车页面
func (r *TaobaoItemCarturlGetAPIRequest) SetCallbackUrl(_callbackUrl string) error {
	r._callbackUrl = _callbackUrl
	r.Set("callback_url", _callbackUrl)
	return nil
}

// GetCallbackUrl CallbackUrl Getter
func (r TaobaoItemCarturlGetAPIRequest) GetCallbackUrl() string {
	return r._callbackUrl
}

// SetUserNick is UserNick Setter
// 商家Nick，优先使用user_id
func (r *TaobaoItemCarturlGetAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaoItemCarturlGetAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetExtParams is ExtParams Setter
// 扩展属性，关注店铺的时候会传递下去，格式为K:V|K:V格式
func (r *TaobaoItemCarturlGetAPIRequest) SetExtParams(_extParams string) error {
	r._extParams = _extParams
	r.Set("ext_params", _extParams)
	return nil
}

// GetExtParams ExtParams Getter
func (r TaobaoItemCarturlGetAPIRequest) GetExtParams() string {
	return r._extParams
}

// SetType is Type Setter
// 端类型，默认是tb，可选tb,hm
func (r *TaobaoItemCarturlGetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoItemCarturlGetAPIRequest) GetType() string {
	return r._type
}

// SetUserId is UserId Setter
// 商家ID
func (r *TaobaoItemCarturlGetAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoItemCarturlGetAPIRequest) GetUserId() int64 {
	return r._userId
}
