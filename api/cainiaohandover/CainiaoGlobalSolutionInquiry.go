package cainiaohandover

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cainiaohandover"
)

/* 
解决方案询盘 
cainiao.global.solution.inquiry

根据交易单号查询可用的解决方案
*/
func CainiaoGlobalSolutionInquiry(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalSolutionInquiryAPIRequest, session string) (*cainiaohandover.CainiaoGlobalSolutionInquiryAPIResponse, error) {
    var resp cainiaohandover.CainiaoGlobalSolutionInquiryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
