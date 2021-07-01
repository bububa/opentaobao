package maitix

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/maitix"
)

/* 
查询招行支付状态api 
alibaba.damai.maitix.distribution.cmb.querypayresult

queryPayResult
*/
func AlibabaDamaiMaitixDistributionCmbQuerypayresult(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIRequest, session string) (*maitix.AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse, error) {
    var resp maitix.AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
