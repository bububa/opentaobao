package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘定向优惠打标接口 API请求
taobao.ump.shoutaotag.add

手淘定向优惠的优惠标签打标接口
给特定的手淘买家打上优惠标记，标记承载在自己的业务标签库中，标签有效期为7天。
*/
type TaobaoUmpShoutaotagAddRequest struct {
    model.Params
    // 商品ID
    _itemId   int64
    // 买家ID
    _buyerId   int64
    // 渠道KEY
    _channelKey   string
}

// 初始化TaobaoUmpShoutaotagAddRequest对象
func NewTaobaoUmpShoutaotagAddRequest() *TaobaoUmpShoutaotagAddRequest{
    return &TaobaoUmpShoutaotagAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpShoutaotagAddRequest) GetApiMethodName() string {
    return "taobao.ump.shoutaotag.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpShoutaotagAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID
func (r *TaobaoUmpShoutaotagAddRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoUmpShoutaotagAddRequest) GetItemId() int64 {
    return r._itemId
}
// BuyerId Setter
// 买家ID
func (r *TaobaoUmpShoutaotagAddRequest) SetBuyerId(_buyerId int64) error {
    r._buyerId = _buyerId
    r.Set("buyer_id", _buyerId)
    return nil
}

// BuyerId Getter
func (r TaobaoUmpShoutaotagAddRequest) GetBuyerId() int64 {
    return r._buyerId
}
// ChannelKey Setter
// 渠道KEY
func (r *TaobaoUmpShoutaotagAddRequest) SetChannelKey(_channelKey string) error {
    r._channelKey = _channelKey
    r.Set("channel_key", _channelKey)
    return nil
}

// ChannelKey Getter
func (r TaobaoUmpShoutaotagAddRequest) GetChannelKey() string {
    return r._channelKey
}
