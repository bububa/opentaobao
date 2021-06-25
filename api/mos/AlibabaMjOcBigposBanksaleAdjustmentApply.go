package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
大pos银行卡调账申请 
alibaba.mj.oc.bigpos.banksale.adjustment.apply

大pos银行卡调账申请
*/
func AlibabaMjOcBigposBanksaleAdjustmentApply(clt *core.SDKClient, req *mos.AlibabaMjOcBigposBanksaleAdjustmentApplyRequest, session string) (*mos.AlibabaMjOcBigposBanksaleAdjustmentApplyResponse, error) {
    var resp mos.AlibabaMjOcBigposBanksaleAdjustmentApplyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
