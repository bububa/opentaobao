package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
产品铺货状态查询 
tmall.supplychain.channel.product.release.status.get

巴拿马战役渠道产品状态查询
*/
func TmallSupplychainChannelProductReleaseStatusGet(clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductReleaseStatusGetRequest, session string) (*fenxiao.TmallSupplychainChannelProductReleaseStatusGetResponse, error) {
    var resp fenxiao.TmallSupplychainChannelProductReleaseStatusGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
