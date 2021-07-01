package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
正大鸡蛋激活追溯码 
alibaba.alihealth.tracecodeplatform.code.active

用于正大鸡蛋激活追溯码
*/
func AlibabaAlihealthTracecodeplatformCodeActive(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodeplatformCodeActiveAPIRequest, session string) (*alihealth2.AlibabaAlihealthTracecodeplatformCodeActiveAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthTracecodeplatformCodeActiveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
