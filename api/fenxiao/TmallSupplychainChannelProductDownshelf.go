package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
产品下架 
tmall.supplychain.channel.product.downshelf

产品下架
*/
func TmallSupplychainChannelProductDownshelf(clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductDownshelfRequest, session string) (*fenxiao.TmallSupplychainChannelProductDownshelfResponse, error) {
    var resp fenxiao.TmallSupplychainChannelProductDownshelfAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
