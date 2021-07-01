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
type TaobaoUmpShoutaotagAddAPIRequest struct {
    model.Params
    // 商品ID
    _itemId   int64
    // 买家ID
    _buyerId   int64
    // 渠道KEY
    _channelKey   string
}

// 初始化TaobaoUmpShoutaotagAddAPIRequest对象
func NewTaobaoUmpShoutaotagAddRequest() *TaobaoUmpShoutaotagAddAPIRequest{
    return &TaobaoUmpShoutaotagAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUmpShoutaotagAddAPIRequest) GetApiMethodName() string {
    return "taobao.ump.shoutaotag.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUmpShoutaotagAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID
func (r *TaobaoUmpShoutaotagAddAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoUmpShoutaotagAddAPIRequest) GetItemId() int64 {
    return r._itemId
}
// BuyerId Setter
// 买家ID
func (r *TaobaoUmpShoutaotagAddAPIRequest) SetBuyerId(_buyerId int64) error {
    r._buyerId = _buyerId
    r.Set("buyer_id", _buyerId)
    return nil
}

// BuyerId Getter
func (r TaobaoUmpShoutaotagAddAPIRequest) GetBuyerId() int64 {
    return r._buyerId
}
// ChannelKey Setter
// 渠道KEY
func (r *TaobaoUmpShoutaotagAddAPIRequest) SetChannelKey(_channelKey string) error {
    r._channelKey = _channelKey
    r.Set("channel_key", _channelKey)
    return nil
}

// ChannelKey Getter
func (r TaobaoUmpShoutaotagAddAPIRequest) GetChannelKey() string {
    return r._channelKey
}
