package legalsuit

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/legalsuit"
)

/* 
更新或者保存管辖信息 
alibaba.legal.suit.domination.push

ISV推送管辖信息到诉讼平台
*/
func AlibabaLegalSuitDominationPush(clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitDominationPushRequest, session string) (*legalsuit.AlibabaLegalSuitDominationPushAPIResponse, error) {
    var resp legalsuit.AlibabaLegalSuitDominationPushAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
