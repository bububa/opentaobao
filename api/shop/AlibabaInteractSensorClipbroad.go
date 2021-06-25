package shop

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/shop"
)

/* 
Weex页面设置或读取剪切板 
alibaba.interact.sensor.clipbroad

Weex页面设置或读取剪切板
*/
func AlibabaInteractSensorClipbroad(clt *core.SDKClient, req *shop.AlibabaInteractSensorClipbroadRequest, session string) (*shop.AlibabaInteractSensorClipbroadResponse, error) {
    var resp shop.AlibabaInteractSensorClipbroadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
