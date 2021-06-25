package icbu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbu"
)

/* 
增加商品分组 
alibaba.icbu.product.group.add

增加商品分组
*/
func AlibabaIcbuProductGroupAdd(clt *core.SDKClient, req *icbu.AlibabaIcbuProductGroupAddRequest, session string) (*icbu.AlibabaIcbuProductGroupAddResponse, error) {
    var resp icbu.AlibabaIcbuProductGroupAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
