package alihouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihouse"
)

/* 
查询楼盘相关信息 
alibaba.alihouse.newhome.project.query

根据outerid查询楼盘相关信息
*/
func AlibabaAlihouseNewhomeProjectQuery(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectQueryRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectQueryAPIResponse, error) {
    var resp alihouse.AlibabaAlihouseNewhomeProjectQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
