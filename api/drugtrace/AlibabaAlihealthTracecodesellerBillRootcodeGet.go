package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
获取最外层包装码 
alibaba.alihealth.tracecodeseller.bill.rootcode.get

获取最外层包装码
*/
func AlibabaAlihealthTracecodesellerBillRootcodeGet(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthTracecodesellerBillRootcodeGetAPIRequest, session string) (*drugtrace.AlibabaAlihealthTracecodesellerBillRootcodeGetAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthTracecodesellerBillRootcodeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
