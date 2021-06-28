package aliqin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliqin"
)

/* 
判断设备是否存在 
alibaba.aliqin.fc.iot.device.isexist

判断设备是否存在
*/
func AlibabaAliqinFcIotDeviceIsexist(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotDeviceIsexistRequest, session string) (*aliqin.AlibabaAliqinFcIotDeviceIsexistAPIResponse, error) {
    var resp aliqin.AlibabaAliqinFcIotDeviceIsexistAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
