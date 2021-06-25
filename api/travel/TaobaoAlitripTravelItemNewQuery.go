package travel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/travel"
)

/* 
【API3.0】新版度假单个商品查询接口 
taobao.alitrip.travel.item.new.query

新版旅行度假新商品查询接口（单个商品查询）
*/
func TaobaoAlitripTravelItemNewQuery(clt *core.SDKClient, req *travel.TaobaoAlitripTravelItemNewQueryRequest, session string) (*travel.TaobaoAlitripTravelItemNewQueryResponse, error) {
    var resp travel.TaobaoAlitripTravelItemNewQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
