package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品优惠详情查询 APIRequest
taobao.ump.promotion.get

商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
*/
type TaobaoUmpPromotionGetRequest struct {
    model.Params

    // 商品id
    itemId   int64 

}

func NewTaobaoUmpPromotionGetRequest() *TaobaoUmpPromotionGetRequest{
    return &TaobaoUmpPromotionGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUmpPromotionGetRequest) GetApiMethodName() string {
    return "taobao.ump.promotion.get"
}

func (r TaobaoUmpPromotionGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUmpPromotionGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoUmpPromotionGetRequest) GetItemId() int64 {
    return r.itemId
}

