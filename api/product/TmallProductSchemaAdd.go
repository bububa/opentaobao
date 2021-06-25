package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
使用Schema文件发布一个产品 
tmall.product.schema.add

Schema体系发布一个产品
*/
func TmallProductSchemaAdd(clt *core.SDKClient, req *product.TmallProductSchemaAddRequest, session string) (*product.TmallProductSchemaAddResponse, error) {
    var resp product.TmallProductSchemaAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
