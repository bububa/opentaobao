package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件上商品优惠获取 APIRequest
tmall.device.item.promotion.get

商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
*/
type TmallDeviceItemPromotionGetRequest struct {
    model.Params

    // 商品id
    itemId   int64 

}

func NewTmallDeviceItemPromotionGetRequest() *TmallDeviceItemPromotionGetRequest{
    return &TmallDeviceItemPromotionGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallDeviceItemPromotionGetRequest) GetApiMethodName() string {
    return "tmall.device.item.promotion.get"
}

func (r TmallDeviceItemPromotionGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallDeviceItemPromotionGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallDeviceItemPromotionGetRequest) GetItemId() int64 {
    return r.itemId
}

