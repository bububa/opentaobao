package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
开放设备id转换内部设备id APIResponse
taobao.ailab.aicloud.top.device.deviceid.convert

将开放设备id转换为内部设备id
*/
type TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAilabAicloudTopDeviceDeviceidConvertResponse `json:"taobao_ailab_aicloud_top_device_deviceid_convert_response,omitempty"`
}

type TaobaoAilabAicloudTopDeviceDeviceidConvertResponse struct {

    // 接口返回model
    Result   *TaobaoAilabAicloudTopDeviceDeviceidConvertResult `json:"result,omitempty"`

}
