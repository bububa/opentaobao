package alihouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihouse"
)

/* 
楼盘上下架 
alibaba.alihouse.newhome.project.line

上下架楼盘
*/
func AlibabaAlihouseNewhomeProjectLine(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectLineRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectLineAPIResponse, error) {
    var resp alihouse.AlibabaAlihouseNewhomeProjectLineAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
