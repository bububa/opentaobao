package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
获取药品扫码落地页vivo 
alibaba.alihealth.tracecodesearch.getshowurl.vivo

获取药品扫码落地页vivo
*/
func AlibabaAlihealthTracecodesearchGetshowurlVivo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthTracecodesearchGetshowurlVivoRequest, session string) (*drugtrace.AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthTracecodesearchGetshowurlVivoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}