package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
渠道价格查询接口 
tmall.supplychain.channel.product.price.get

渠道价格查询接口
*/
func TmallSupplychainChannelProductPriceGet(clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductPriceGetRequest, session string) (*fenxiao.TmallSupplychainChannelProductPriceGetResponse, error) {
    var resp fenxiao.TmallSupplychainChannelProductPriceGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
