package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
美妆虚拟试装 
alibaba.interact.sensor.makeup

手机淘宝美妆类目虚拟试妆权限，客户端能力（JS－API）
*/
func AlibabaInteractSensorMakeup(clt *core.SDKClient, req *interact.AlibabaInteractSensorMakeupRequest, session string) (*interact.AlibabaInteractSensorMakeupAPIResponse, error) {
    var resp interact.AlibabaInteractSensorMakeupAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
