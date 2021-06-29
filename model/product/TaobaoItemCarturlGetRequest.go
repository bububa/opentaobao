package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
加购URL获取 API请求
taobao.item.carturl.get

获取加购URL，支持添加商品到购物车
*/
type TaobaoItemCarturlGetRequest struct {
    model.Params
    // 商品信息，格式为 商品ID_SKU ID_数量，多条记录以逗号(,)分割
    _itemIds   []string
    // 回调地址，需要是EWS域名地址。可不填，默认到购物车页面
    _callbackUrl   string
    // 商家Nick，优先使用user_id
    _userNick   string
    // 商家ID
    _userId   int64
    // 扩展属性，关注店铺的时候会传递下去，格式为K:V|K:V格式
    _extParams   string
    // 端类型，默认是tb，可选tb,hm
    _type   string
}

// 初始化TaobaoItemCarturlGetRequest对象
func NewTaobaoItemCarturlGetRequest() *TaobaoItemCarturlGetRequest{
    return &TaobaoItemCarturlGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemCarturlGetRequest) GetApiMethodName() string {
    return "taobao.item.carturl.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemCarturlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemIds Setter
// 商品信息，格式为 商品ID_SKU ID_数量，多条记录以逗号(,)分割
func (r *TaobaoItemCarturlGetRequest) SetItemIds(_itemIds []string) error {
    r._itemIds = _itemIds
    r.Set("item_ids", _itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoItemCarturlGetRequest) GetItemIds() []string {
    return r._itemIds
}
// CallbackUrl Setter
// 回调地址，需要是EWS域名地址。可不填，默认到购物车页面
func (r *TaobaoItemCarturlGetRequest) SetCallbackUrl(_callbackUrl string) error {
    r._callbackUrl = _callbackUrl
    r.Set("callback_url", _callbackUrl)
    return nil
}

// CallbackUrl Getter
func (r TaobaoItemCarturlGetRequest) GetCallbackUrl() string {
    return r._callbackUrl
}
// UserNick Setter
// 商家Nick，优先使用user_id
func (r *TaobaoItemCarturlGetRequest) SetUserNick(_userNick string) error {
    r._userNick = _userNick
    r.Set("user_nick", _userNick)
    return nil
}

// UserNick Getter
func (r TaobaoItemCarturlGetRequest) GetUserNick() string {
    return r._userNick
}
// UserId Setter
// 商家ID
func (r *TaobaoItemCarturlGetRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoItemCarturlGetRequest) GetUserId() int64 {
    return r._userId
}
// ExtParams Setter
// 扩展属性，关注店铺的时候会传递下去，格式为K:V|K:V格式
func (r *TaobaoItemCarturlGetRequest) SetExtParams(_extParams string) error {
    r._extParams = _extParams
    r.Set("ext_params", _extParams)
    return nil
}

// ExtParams Getter
func (r TaobaoItemCarturlGetRequest) GetExtParams() string {
    return r._extParams
}
// Type Setter
// 端类型，默认是tb，可选tb,hm
func (r *TaobaoItemCarturlGetRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoItemCarturlGetRequest) GetType() string {
    return r._type
}
