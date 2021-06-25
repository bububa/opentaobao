package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假线路商品上下架接口 APIRequest
taobao.alitrip.travel.item.shelve

旅行度假新商品发布接口 第三版：度假商品上下架接口
注意：定时上下架功能，目前只支持接送、租车类目。
*/
type TaobaoAlitripTravelItemShelveRequest struct {
    model.Params

    // 商品id。itemId和outProductId至少填写一个
    itemId   int64 

    // 商品 外部商家编码。itemId和outProductId至少填写一个
    outProductId   string 

    // 1-上架 0-下架
    itemStatus   int64 

    // 指定定时上架时间，格式：yyyy-MM-dd HH:mm:ss。若不设置该值且item_status为1，则表示立即上架。
    onlineTime   string 

}

func NewTaobaoAlitripTravelItemShelveRequest() *TaobaoAlitripTravelItemShelveRequest{
    return &TaobaoAlitripTravelItemShelveRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelItemShelveRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.item.shelve"
}

func (r TaobaoAlitripTravelItemShelveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelItemShelveRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoAlitripTravelItemShelveRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoAlitripTravelItemShelveRequest) SetOutProductId(outProductId string) error {
    r.outProductId = outProductId
    r.Set("out_product_id", outProductId)
    return nil
}

func (r TaobaoAlitripTravelItemShelveRequest) GetOutProductId() string {
    return r.outProductId
}

func (r *TaobaoAlitripTravelItemShelveRequest) SetItemStatus(itemStatus int64) error {
    r.itemStatus = itemStatus
    r.Set("item_status", itemStatus)
    return nil
}

func (r TaobaoAlitripTravelItemShelveRequest) GetItemStatus() int64 {
    return r.itemStatus
}

func (r *TaobaoAlitripTravelItemShelveRequest) SetOnlineTime(onlineTime string) error {
    r.onlineTime = onlineTime
    r.Set("online_time", onlineTime)
    return nil
}

func (r TaobaoAlitripTravelItemShelveRequest) GetOnlineTime() string {
    return r.onlineTime
}

