package legalcase

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/legalcase"
)

/* 
新增反馈口径 
alibaba.legal.case.standpoint.savestandpoint

新增反馈口径 ,从外部接受反馈的口径
*/
func AlibabaLegalCaseStandpointSavestandpoint(clt *core.SDKClient, req *legalcase.AlibabaLegalCaseStandpointSavestandpointRequest, session string) (*legalcase.AlibabaLegalCaseStandpointSavestandpointResponse, error) {
    var resp legalcase.AlibabaLegalCaseStandpointSavestandpointAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
