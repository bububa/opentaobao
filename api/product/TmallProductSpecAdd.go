package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
添加产品规格 
tmall.product.spec.add

增加产品规格
*/
func TmallProductSpecAdd(clt *core.SDKClient, req *product.TmallProductSpecAddRequest, session string) (*product.TmallProductSpecAddResponse, error) {
    var resp product.TmallProductSpecAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
