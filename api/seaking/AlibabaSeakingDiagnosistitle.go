package seaking

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/seaking"
)

/* 
标题诊断 
alibaba.seaking.diagnosistitle

标题诊断
*/
func AlibabaSeakingDiagnosistitle(clt *core.SDKClient, req *seaking.AlibabaSeakingDiagnosistitleAPIRequest, session string) (*seaking.AlibabaSeakingDiagnosistitleAPIResponse, error) {
    var resp seaking.AlibabaSeakingDiagnosistitleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
