package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
加购URL获取 APIRequest
taobao.item.carturl.get

获取加购URL，支持添加商品到购物车
*/
type TaobaoItemCarturlGetRequest struct {
    model.Params

    // 商品信息，格式为 商品ID_SKU ID_数量，多条记录以逗号(,)分割
    itemIds   []String 

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

func NewTaobaoItemCarturlGetRequest() *TaobaoItemCarturlGetRequest{
    return &TaobaoItemCarturlGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemCarturlGetRequest) GetApiMethodName() string {
    return "taobao.item.carturl.get"
}

func (r TaobaoItemCarturlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemCarturlGetRequest) SetItemIds(itemIds []String) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

func (r TaobaoItemCarturlGetRequest) GetItemIds() []String {
    return r.itemIds
}

func (r *TaobaoItemCarturlGetRequest) SetCallbackUrl(callbackUrl string) error {
    r.callbackUrl = callbackUrl
    r.Set("callback_url", callbackUrl)
    return nil
}

func (r TaobaoItemCarturlGetRequest) GetCallbackUrl() string {
    return r.callbackUrl
}

func (r *TaobaoItemCarturlGetRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

func (r TaobaoItemCarturlGetRequest) GetUserNick() string {
    return r.userNick
}

func (r *TaobaoItemCarturlGetRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoItemCarturlGetRequest) GetUserId() int64 {
    return r.userId
}

func (r *TaobaoItemCarturlGetRequest) SetExtParams(extParams string) error {
    r.extParams = extParams
    r.Set("ext_params", extParams)
    return nil
}

func (r TaobaoItemCarturlGetRequest) GetExtParams() string {
    return r.extParams
}

func (r *TaobaoItemCarturlGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoItemCarturlGetRequest) GetType() string {
    return r.type
}

