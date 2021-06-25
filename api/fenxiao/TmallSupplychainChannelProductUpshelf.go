package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
产品上架 
tmall.supplychain.channel.product.upshelf

上架渠道产品
*/
func TmallSupplychainChannelProductUpshelf(clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductUpshelfRequest, session string) (*fenxiao.TmallSupplychainChannelProductUpshelfResponse, error) {
    var resp fenxiao.TmallSupplychainChannelProductUpshelfAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
