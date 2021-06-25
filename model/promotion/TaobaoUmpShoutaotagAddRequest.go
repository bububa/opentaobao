package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘定向优惠打标接口 APIRequest
taobao.ump.shoutaotag.add

手淘定向优惠的优惠标签打标接口
给特定的手淘买家打上优惠标记，标记承载在自己的业务标签库中，标签有效期为7天。
*/
type TaobaoUmpShoutaotagAddRequest struct {
    model.Params

    // 商品ID
    itemId   int64 

    // 买家ID
    buyerId   int64 

    // 渠道KEY
    channelKey   string 

}

func NewTaobaoUmpShoutaotagAddRequest() *TaobaoUmpShoutaotagAddRequest{
    return &TaobaoUmpShoutaotagAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpShoutaotagAddRequest) GetApiMethodName() string {
    return "taobao.ump.shoutaotag.add"
}

func (r TaobaoUmpShoutaotagAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpShoutaotagAddRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoUmpShoutaotagAddRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoUmpShoutaotagAddRequest) SetBuyerId(buyerId int64) error {
    r.buyerId = buyerId
    r.Set("buyer_id", buyerId)
    return nil
}

func (r TaobaoUmpShoutaotagAddRequest) GetBuyerId() int64 {
    return r.buyerId
}

func (r *TaobaoUmpShoutaotagAddRequest) SetChannelKey(channelKey string) error {
    r.channelKey = channelKey
    r.Set("channel_key", channelKey)
    return nil
}

func (r TaobaoUmpShoutaotagAddRequest) GetChannelKey() string {
    return r.channelKey
}

