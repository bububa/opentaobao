package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】新版度假单个商品查询接口 API请求
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

// 初始化TaobaoAlitripTravelItemNewQueryRequest对象
func NewTaobaoAlitripTravelItemNewQueryRequest() *TaobaoAlitripTravelItemNewQueryRequest{
    return &TaobaoAlitripTravelItemNewQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemNewQueryRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.item.new.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemNewQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemNewQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoAlitripTravelItemNewQueryRequest) GetItemId() int64 {
    return r.itemId
}
// OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemNewQueryRequest) SetOutProductId(outProductId string) error {
    r.outProductId = outProductId
    r.Set("out_product_id", outProductId)
    return nil
}

// OutProductId Getter
func (r TaobaoAlitripTravelItemNewQueryRequest) GetOutProductId() string {
    return r.outProductId
}
