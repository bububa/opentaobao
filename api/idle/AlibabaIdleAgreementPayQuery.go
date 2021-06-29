package idle

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idle"
)

/* 
代扣详情查询 
alibaba.idle.agreement.pay.query

查询代扣结果
*/
func AlibabaIdleAgreementPayQuery(clt *core.SDKClient, req *idle.AlibabaIdleAgreementPayQueryRequest, session string) (*idle.AlibabaIdleAgreementPayQueryAPIResponse, error) {
    var resp idle.AlibabaIdleAgreementPayQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
