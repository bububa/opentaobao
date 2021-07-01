package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
区域价格查询 
taobao.region.price.query

区域价格查询
*/
func TaobaoRegionPriceQuery(clt *core.SDKClient, req *fenxiao.TaobaoRegionPriceQueryAPIRequest, session string) (*fenxiao.TaobaoRegionPriceQueryAPIResponse, error) {
    var resp fenxiao.TaobaoRegionPriceQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
