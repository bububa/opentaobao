package legalcase

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/legalcase"
)

/* 
委托 
alibaba.legal.case.entrust.get

获取委托案件的基本信息
*/
func AlibabaLegalCaseEntrustGet(clt *core.SDKClient, req *legalcase.AlibabaLegalCaseEntrustGetRequest, session string) (*legalcase.AlibabaLegalCaseEntrustGetAPIResponse, error) {
    var resp legalcase.AlibabaLegalCaseEntrustGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
