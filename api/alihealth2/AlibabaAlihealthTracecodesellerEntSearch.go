package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
查询商家信息 
alibaba.alihealth.tracecodeseller.ent.search

查询商家信息
*/
func AlibabaAlihealthTracecodesellerEntSearch(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerEntSearchRequest, session string) (*alihealth2.AlibabaAlihealthTracecodesellerEntSearchAPIResponse, error) {
    var resp alihealth2.AlibabaAlihealthTracecodesellerEntSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
