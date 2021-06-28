package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
获取产品的规格信息 
tmall.product.specs.get

按product_id或品牌下载产品规格，返回一组的产品规格信息。
*/
func TmallProductSpecsGet(clt *core.SDKClient, req *product.TmallProductSpecsGetRequest, session string) (*product.TmallProductSpecsGetAPIResponse, error) {
    var resp product.TmallProductSpecsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
