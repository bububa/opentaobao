package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
算法获取hscode 
tmall.item.calculate.hscode.get

算法获取hscode
*/
func TmallItemCalculateHscodeGet(clt *core.SDKClient, req *product.TmallItemCalculateHscodeGetRequest, session string) (*product.TmallItemCalculateHscodeGetResponse, error) {
    var resp product.TmallItemCalculateHscodeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
