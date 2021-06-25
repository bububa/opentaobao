package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
基本ui操作 
alibaba.interact.sensor.ui

Weex 基本UI操作
*/
func AlibabaInteractSensorUi(clt *core.SDKClient, req *util.AlibabaInteractSensorUiRequest, session string) (*util.AlibabaInteractSensorUiResponse, error) {
    var resp util.AlibabaInteractSensorUiAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
