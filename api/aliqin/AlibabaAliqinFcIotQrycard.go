package aliqin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliqin"
)

/* 
查询终端信息 
alibaba.aliqin.fc.iot.qrycard

查询终端信息
*/
func AlibabaAliqinFcIotQrycard(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotQrycardRequest, session string) (*aliqin.AlibabaAliqinFcIotQrycardResponse, error) {
    var resp aliqin.AlibabaAliqinFcIotQrycardAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
