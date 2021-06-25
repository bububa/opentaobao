package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
产品信息获取schema获取 
tmall.product.schema.get

产品信息获取接口schema形式返回
*/
func TmallProductSchemaGet(clt *core.SDKClient, req *product.TmallProductSchemaGetRequest, session string) (*product.TmallProductSchemaGetResponse, error) {
    var resp product.TmallProductSchemaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
