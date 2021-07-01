package tmallnr

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallnr"
)

/* 
同步商户中心服务范围 
tmall.nr.seller.storerange.sync

同步商户中心服务范围
*/
func TmallNrSellerStorerangeSync(clt *core.SDKClient, req *tmallnr.TmallNrSellerStorerangeSyncAPIRequest, session string) (*tmallnr.TmallNrSellerStorerangeSyncAPIResponse, error) {
    var resp tmallnr.TmallNrSellerStorerangeSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
