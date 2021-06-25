package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
淘口令检查 
alibaba.baichuan.taopassword.check

检查当前文本是否为淘口令
*/
func AlibabaBaichuanTaopasswordCheck(clt *core.SDKClient, req *baichuan.AlibabaBaichuanTaopasswordCheckRequest, session string) (*baichuan.AlibabaBaichuanTaopasswordCheckResponse, error) {
    var resp baichuan.AlibabaBaichuanTaopasswordCheckAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
