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
    itemIds   []string
    // 回调地址，需要是EWS域名地址。可不填，默认到购物车页面
    callbackUrl   string
    // 商家Nick，优先使用user_id
    userNick   string
    // 商家ID
    userId   int64
    // 扩展属性，关注店铺的时候会传递下去，格式为K:V|K:V格式
    extParams   string
    // 端类型，默认是tb，可选tb,hm
    type   string
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
func (r *TaobaoItemCarturlGetRequest) SetItemIds(itemIds []string) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoItemCarturlGetRequest) GetItemIds() []string {
    return r.itemIds
}
// CallbackUrl Setter
// 回调地址，需要是EWS域名地址。可不填，默认到购物车页面
func (r *TaobaoItemCarturlGetRequest) SetCallbackUrl(callbackUrl string) error {
    r.callbackUrl = callbackUrl
    r.Set("callback_url", callbackUrl)
    return nil
}

// CallbackUrl Getter
func (r TaobaoItemCarturlGetRequest) GetCallbackUrl() string {
    return r.callbackUrl
}
// UserNick Setter
// 商家Nick，优先使用user_id
func (r *TaobaoItemCarturlGetRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

// UserNick Getter
func (r TaobaoItemCarturlGetRequest) GetUserNick() string {
    return r.userNick
}
// UserId Setter
// 商家ID
func (r *TaobaoItemCarturlGetRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoItemCarturlGetRequest) GetUserId() int64 {
    return r.userId
}
// ExtParams Setter
// 扩展属性，关注店铺的时候会传递下去，格式为K:V|K:V格式
func (r *TaobaoItemCarturlGetRequest) SetExtParams(extParams string) error {
    r.extParams = extParams
    r.Set("ext_params", extParams)
    return nil
}

// ExtParams Getter
func (r TaobaoItemCarturlGetRequest) GetExtParams() string {
    return r.extParams
}
// Type Setter
// 端类型，默认是tb，可选tb,hm
func (r *TaobaoItemCarturlGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoItemCarturlGetRequest) GetType() string {
    return r.type
}
