package legalsuit

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/legalsuit"
)

/* 
委托开庭服务查询 
alibaba.legal.suit.court.entrust.get

查询委托开庭信息
*/
func AlibabaLegalSuitCourtEntrustGet(clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitCourtEntrustGetRequest, session string) (*legalsuit.AlibabaLegalSuitCourtEntrustGetAPIResponse, error) {
    var resp legalsuit.AlibabaLegalSuitCourtEntrustGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
