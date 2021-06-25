package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
homs人力资源组织树查询接口 
alibaba.wdk.hrworkbench.cdporgs.query

提供查询homs人力组织树的接口，按照商家做权限隔离。
*/
func AlibabaWdkHrworkbenchCdporgsQuery(clt *core.SDKClient, req *wdk.AlibabaWdkHrworkbenchCdporgsQueryRequest, session string) (*wdk.AlibabaWdkHrworkbenchCdporgsQueryResponse, error) {
    var resp wdk.AlibabaWdkHrworkbenchCdporgsQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
