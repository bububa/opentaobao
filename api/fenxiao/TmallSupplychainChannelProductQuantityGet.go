package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
渠道库存查询接口 
tmall.supplychain.channel.product.quantity.get

渠道库存查询接口
*/
func TmallSupplychainChannelProductQuantityGet(clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductQuantityGetRequest, session string) (*fenxiao.TmallSupplychainChannelProductQuantityGetAPIResponse, error) {
    var resp fenxiao.TmallSupplychainChannelProductQuantityGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
