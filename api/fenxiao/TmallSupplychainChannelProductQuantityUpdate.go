package fenxiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fenxiao"
)

/* 
渠道无仓库存更新接口 
tmall.supplychain.channel.product.quantity.update

渠道无仓库存更新接口
*/
func TmallSupplychainChannelProductQuantityUpdate(clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductQuantityUpdateRequest, session string) (*fenxiao.TmallSupplychainChannelProductQuantityUpdateResponse, error) {
    var resp fenxiao.TmallSupplychainChannelProductQuantityUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
