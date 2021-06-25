package interact

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/interact"
)

/* 
获取登陆页面 
alibaba.interact.sensor.login

获取登陆页面
*/
func AlibabaInteractSensorLogin(clt *core.SDKClient, req *interact.AlibabaInteractSensorLoginRequest, session string) (*interact.AlibabaInteractSensorLoginResponse, error) {
    var resp interact.AlibabaInteractSensorLoginAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
