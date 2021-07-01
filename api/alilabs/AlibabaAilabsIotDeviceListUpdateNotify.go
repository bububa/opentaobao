package alilabs

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alilabs"
)

/* 
设备列表更新通知 
alibaba.ailabs.iot.device.list.update.notify

用于人工智能实验室IoT合作厂商上报三方接入IoT设备列表更新时的通知
*/
func AlibabaAilabsIotDeviceListUpdateNotify(clt *core.SDKClient, req *alilabs.AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest, session string) (*alilabs.AlibabaAilabsIotDeviceListUpdateNotifyAPIResponse, error) {
    var resp alilabs.AlibabaAilabsIotDeviceListUpdateNotifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
