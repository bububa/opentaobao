package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
获取vivo banner 
alibaba.alihealth.tracecodesearc.getinfomation.vivo

获取vivo banner  url
*/
func AlibabaAlihealthTracecodesearcGetinfomationVivo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthTracecodesearcGetinfomationVivoRequest, session string) (*drugtrace.AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthTracecodesearcGetinfomationVivoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}