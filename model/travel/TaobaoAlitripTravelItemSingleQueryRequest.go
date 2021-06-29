package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假单个商品查询接口 API请求
taobao.alitrip.travel.item.single.query

旅行度假新商品查询接口（单个商品查询） 第三版
*/
type TaobaoAlitripTravelItemSingleQueryRequest struct {
    model.Params
    // 商品id。itemId和outProductId至少填写一个
    itemId   int64
    // 商品 外部商家编码。itemId和outProductId至少填写一个
    outProductId   string
}

// 初始化TaobaoAlitripTravelItemSingleQueryRequest对象
func NewTaobaoAlitripTravelItemSingleQueryRequest() *TaobaoAlitripTravelItemSingleQueryRequest{
    return &TaobaoAlitripTravelItemSingleQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemSingleQueryRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.item.single.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemSingleQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSingleQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoAlitripTravelItemSingleQueryRequest) GetItemId() int64 {
    return r.itemId
}
// OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSingleQueryRequest) SetOutProductId(outProductId string) error {
    r.outProductId = outProductId
    r.Set("out_product_id", outProductId)
    return nil
}

// OutProductId Getter
func (r TaobaoAlitripTravelItemSingleQueryRequest) GetOutProductId() string {
    return r.outProductId
}
