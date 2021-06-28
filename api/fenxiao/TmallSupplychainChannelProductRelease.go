package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
供应商铺货 
tmall.supplychain.channel.product.release

供应商渠道铺货接口
*/
func TmallSupplychainChannelProductRelease(clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductReleaseRequest, session string) (*fenxiao.TmallSupplychainChannelProductReleaseAPIResponse, error) {
    var resp fenxiao.TmallSupplychainChannelProductReleaseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
