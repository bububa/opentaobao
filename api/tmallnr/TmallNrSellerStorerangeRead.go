package tmallnr

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallnr"
)

/* 
门店服务范围读取 
tmall.nr.seller.storerange.read

读取卖家所属门店的服务范围
*/
func TmallNrSellerStorerangeRead(clt *core.SDKClient, req *tmallnr.TmallNrSellerStorerangeReadRequest, session string) (*tmallnr.TmallNrSellerStorerangeReadAPIResponse, error) {
    var resp tmallnr.TmallNrSellerStorerangeReadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
