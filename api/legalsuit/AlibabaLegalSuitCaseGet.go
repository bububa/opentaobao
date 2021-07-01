package legalsuit

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/legalsuit"
)

/* 
获取案件信息接口v2版本 
alibaba.legal.suit.case.get

获取案件信息
*/
func AlibabaLegalSuitCaseGet(clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitCaseGetAPIRequest, session string) (*legalsuit.AlibabaLegalSuitCaseGetAPIResponse, error) {
    var resp legalsuit.AlibabaLegalSuitCaseGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
