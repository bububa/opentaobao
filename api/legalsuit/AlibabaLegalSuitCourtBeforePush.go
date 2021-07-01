package legalsuit

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/legalsuit"
)

/* 
更新或保存庭前信息 
alibaba.legal.suit.court.before.push

更新或者保存庭前信息
*/
func AlibabaLegalSuitCourtBeforePush(clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitCourtBeforePushAPIRequest, session string) (*legalsuit.AlibabaLegalSuitCourtBeforePushAPIResponse, error) {
    var resp legalsuit.AlibabaLegalSuitCourtBeforePushAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
