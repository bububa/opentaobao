package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
百川支付完成回调 
taobao.baichuan.payresult.query

百川支付完成回调
*/
func TaobaoBaichuanPayresultQuery(clt *core.SDKClient, req *baichuan.TaobaoBaichuanPayresultQueryAPIRequest, session string) (*baichuan.TaobaoBaichuanPayresultQueryAPIResponse, error) {
    var resp baichuan.TaobaoBaichuanPayresultQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
