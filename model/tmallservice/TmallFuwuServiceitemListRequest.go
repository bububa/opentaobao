package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取服务商品扩展信息 APIRequest
tmall.fuwu.serviceitem.list

获取服务商品扩展信息
*/
type TmallFuwuServiceitemListRequest struct {
    model.Params

    // 商品所属卖家账号id
    sellerId   int64 

    // 商品id列表，有数量限制
    itemids   []int64 

}

func NewTmallFuwuServiceitemListRequest() *TmallFuwuServiceitemListRequest{
    return &TmallFuwuServiceitemListRequest{
        Params: model.NewParams(),
    }
}

func (r TmallFuwuServiceitemListRequest) GetApiMethodName() string {
    return "tmall.fuwu.serviceitem.list"
}

func (r TmallFuwuServiceitemListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallFuwuServiceitemListRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

func (r TmallFuwuServiceitemListRequest) GetSellerId() int64 {
    return r.sellerId
}

func (r *TmallFuwuServiceitemListRequest) SetItemids(itemids []int64) error {
    r.itemids = itemids
    r.Set("itemids", itemids)
    return nil
}

func (r TmallFuwuServiceitemListRequest) GetItemids() []int64 {
    return r.itemids
}

