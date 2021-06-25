package mos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mos"
)

/* 
云屏埋点数据记录接口 
alibaba.mos.store.recordscreenpointinfo

记录云屏埋点数据
*/
func AlibabaMosStoreRecordscreenpointinfo(clt *core.SDKClient, req *mos.AlibabaMosStoreRecordscreenpointinfoRequest, session string) (*mos.AlibabaMosStoreRecordscreenpointinfoResponse, error) {
    var resp mos.AlibabaMosStoreRecordscreenpointinfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
