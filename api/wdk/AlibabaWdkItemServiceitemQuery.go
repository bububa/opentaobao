package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
查询服务商品 
alibaba.wdk.item.serviceitem.query

查询服务商品
*/
func AlibabaWdkItemServiceitemQuery(clt *core.SDKClient, req *wdk.AlibabaWdkItemServiceitemQueryRequest, session string) (*wdk.AlibabaWdkItemServiceitemQueryAPIResponse, error) {
    var resp wdk.AlibabaWdkItemServiceitemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
