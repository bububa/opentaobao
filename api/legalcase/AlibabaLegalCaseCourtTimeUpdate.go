package legalcase

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/legalcase"
)

/* 
开庭时间变更 
alibaba.legal.case.court.time.update

修改案件的开庭时间
*/
func AlibabaLegalCaseCourtTimeUpdate(clt *core.SDKClient, req *legalcase.AlibabaLegalCaseCourtTimeUpdateRequest, session string) (*legalcase.AlibabaLegalCaseCourtTimeUpdateResponse, error) {
    var resp legalcase.AlibabaLegalCaseCourtTimeUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
