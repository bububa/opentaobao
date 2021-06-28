package icbu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbu"
)

/* 
（新）ICBU类目树获取接口 
alibaba.icbu.category.get.new

获取商品发布类目
*/
func AlibabaIcbuCategoryGetNew(clt *core.SDKClient, req *icbu.AlibabaIcbuCategoryGetNewRequest, session string) (*icbu.AlibabaIcbuCategoryGetNewAPIResponse, error) {
    var resp icbu.AlibabaIcbuCategoryGetNewAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
