package maitix

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/maitix"
)

/* 
加密招商一网能支付入参 
alibaba.damai.maitix.distribution.cmb.paramencrypt

encryptParam4Cmb
*/
func AlibabaDamaiMaitixDistributionCmbParamencrypt(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixDistributionCmbParamencryptRequest, session string) (*maitix.AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse, error) {
    var resp maitix.AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
