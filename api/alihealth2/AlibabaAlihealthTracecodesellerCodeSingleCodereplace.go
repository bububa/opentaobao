package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
非药单码替换 
alibaba.alihealth.tracecodeseller.code.single.codereplace

提供非药追溯码单码替换功能
*/
func AlibabaAlihealthTracecodesellerCodeSingleCodereplace(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest, session string) (*alihealth2.AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
