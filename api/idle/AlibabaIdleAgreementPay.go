package idle

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idle"
)

/* 
闲鱼平台商户代扣 
alibaba.idle.agreement.pay

闲鱼平台代扣能力：用户和闲鱼签约代扣协议 服务商通过直付通产品挂载成为闲鱼二级商户 来完成用户和服务商的结算
*/
func AlibabaIdleAgreementPay(clt *core.SDKClient, req *idle.AlibabaIdleAgreementPayRequest, session string) (*idle.AlibabaIdleAgreementPayAPIResponse, error) {
    var resp idle.AlibabaIdleAgreementPayAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
