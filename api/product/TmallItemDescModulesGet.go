package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
商品描述模块信息获取 
tmall.item.desc.modules.get

商品描述模块信息获取，包括运营设定的类目级别的模块信息以及用户自定义模块数量约束。
*/
func TmallItemDescModulesGet(clt *core.SDKClient, req *product.TmallItemDescModulesGetRequest, session string) (*product.TmallItemDescModulesGetResponse, error) {
    var resp product.TmallItemDescModulesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
