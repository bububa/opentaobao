package alihouse

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihouse"
)

/* 
图文草稿状态更新 
alibaba.alihouse.newhome.rc.changestatus

图文草稿状态更新
*/
func AlibabaAlihouseNewhomeRcChangestatus(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeRcChangestatusRequest, session string) (*alihouse.AlibabaAlihouseNewhomeRcChangestatusAPIResponse, error) {
    var resp alihouse.AlibabaAlihouseNewhomeRcChangestatusAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
