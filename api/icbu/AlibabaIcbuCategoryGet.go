package icbu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbu"
)

/* 
商品发布类目获取 
alibaba.icbu.category.get

获取商品发布类目
*/
func AlibabaIcbuCategoryGet(clt *core.SDKClient, req *icbu.AlibabaIcbuCategoryGetAPIRequest, session string) (*icbu.AlibabaIcbuCategoryGetAPIResponse, error) {
    var resp icbu.AlibabaIcbuCategoryGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
