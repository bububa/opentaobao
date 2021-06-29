package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件上商品优惠获取 API请求
tmall.device.item.promotion.get

商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
*/
type TmallDeviceItemPromotionGetRequest struct {
    model.Params
    // 商品id
    _itemId   int64
}

// 初始化TmallDeviceItemPromotionGetRequest对象
func NewTmallDeviceItemPromotionGetRequest() *TmallDeviceItemPromotionGetRequest{
    return &TmallDeviceItemPromotionGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallDeviceItemPromotionGetRequest) GetApiMethodName() string {
    return "tmall.device.item.promotion.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallDeviceItemPromotionGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TmallDeviceItemPromotionGetRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallDeviceItemPromotionGetRequest) GetItemId() int64 {
    return r._itemId
}
