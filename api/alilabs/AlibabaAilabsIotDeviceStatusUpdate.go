package alilabs

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alilabs"
)

/* 
ailabs iot 设备状态更新 
alibaba.ailabs.iot.device.status.update

用于人工智能实验室IoT合作厂商上报三方接入IoT设备状态更新时的设备状态上报
*/
func AlibabaAilabsIotDeviceStatusUpdate(clt *core.SDKClient, req *alilabs.AlibabaAilabsIotDeviceStatusUpdateRequest, session string) (*alilabs.AlibabaAilabsIotDeviceStatusUpdateAPIResponse, error) {
    var resp alilabs.AlibabaAilabsIotDeviceStatusUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
