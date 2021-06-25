package icbu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbu"
)

/* 
分组信息获取 
alibaba.icbu.product.group.get

分组信息获取
*/
func AlibabaIcbuProductGroupGet(clt *core.SDKClient, req *icbu.AlibabaIcbuProductGroupGetRequest, session string) (*icbu.AlibabaIcbuProductGroupGetResponse, error) {
    var resp icbu.AlibabaIcbuProductGroupGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
