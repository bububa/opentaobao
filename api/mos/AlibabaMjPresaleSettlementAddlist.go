package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
预售结算数据回传 
alibaba.mj.presale.settlement.addlist

用于预售活动结算数据的回传。
*/
func AlibabaMjPresaleSettlementAddlist(clt *core.SDKClient, req *mos.AlibabaMjPresaleSettlementAddlistRequest, session string) (*mos.AlibabaMjPresaleSettlementAddlistResponse, error) {
    var resp mos.AlibabaMjPresaleSettlementAddlistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
