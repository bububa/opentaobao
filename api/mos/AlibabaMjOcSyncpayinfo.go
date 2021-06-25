package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
支付参考号回传 
alibaba.mj.oc.syncpayinfo

支付参考号同步到oc
*/
func AlibabaMjOcSyncpayinfo(clt *core.SDKClient, req *mos.AlibabaMjOcSyncpayinfoRequest, session string) (*mos.AlibabaMjOcSyncpayinfoResponse, error) {
    var resp mos.AlibabaMjOcSyncpayinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
