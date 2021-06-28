package aliqin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliqin"
)

/* 
商家提交设备信息 
alibaba.aliqin.fc.iot.device.post

物联网商家设备信息录入
*/
func AlibabaAliqinFcIotDevicePost(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotDevicePostRequest, session string) (*aliqin.AlibabaAliqinFcIotDevicePostAPIResponse, error) {
    var resp aliqin.AlibabaAliqinFcIotDevicePostAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
