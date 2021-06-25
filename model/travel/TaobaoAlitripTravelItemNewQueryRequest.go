package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】新版度假单个商品查询接口 APIRequest
taobao.alitrip.travel.item.new.query

新版旅行度假新商品查询接口（单个商品查询）
*/
type TaobaoAlitripTravelItemNewQueryRequest struct {
    model.Params

    // 商品id。itemId和outProductId至少填写一个
    itemId   int64 

    // 商品 外部商家编码。itemId和outProductId至少填写一个
    outProductId   string 

}

func NewTaobaoAlitripTravelItemNewQueryRequest() *TaobaoAlitripTravelItemNewQueryRequest{
    return &TaobaoAlitripTravelItemNewQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTravelItemNewQueryRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.item.new.query"
}

func (r TaobaoAlitripTravelItemNewQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTravelItemNewQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoAlitripTravelItemNewQueryRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoAlitripTravelItemNewQueryRequest) SetOutProductId(outProductId string) error {
    r.outProductId = outProductId
    r.Set("out_product_id", outProductId)
    return nil
}

func (r TaobaoAlitripTravelItemNewQueryRequest) GetOutProductId() string {
    return r.outProductId
}

