package baoxian

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baoxian"
)

/* 
更新赔案 
alipay.baoxian.claim.update

更新保险理赔单
*/
func AlipayBaoxianClaimUpdate(clt *core.SDKClient, req *baoxian.AlipayBaoxianClaimUpdateAPIRequest, session string) (*baoxian.AlipayBaoxianClaimUpdateAPIResponse, error) {
    var resp baoxian.AlipayBaoxianClaimUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
