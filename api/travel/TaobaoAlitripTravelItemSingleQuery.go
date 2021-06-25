package travel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/travel"
)

/* 
【API3.0】度假单个商品查询接口 
taobao.alitrip.travel.item.single.query

旅行度假新商品查询接口（单个商品查询） 第三版
*/
func TaobaoAlitripTravelItemSingleQuery(clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemSingleQueryRequest, session string) (*travel.TaobaoAlitripTravelItemSingleQueryResponse, error) {
    var resp travel.TaobaoAlitripTravelItemSingleQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
