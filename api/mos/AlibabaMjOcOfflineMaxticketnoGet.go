package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
pos机获取线下最大小票号 
alibaba.mj.oc.offline.maxticketno.get

给pos机提供线下最大小票号查询
*/
func AlibabaMjOcOfflineMaxticketnoGet(clt *core.SDKClient, req *mos.AlibabaMjOcOfflineMaxticketnoGetRequest, session string) (*mos.AlibabaMjOcOfflineMaxticketnoGetResponse, error) {
    var resp mos.AlibabaMjOcOfflineMaxticketnoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
