package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备扩展信息 API返回值 
taobao.ailab.aicloud.top.device.extinfo.get

获取设备扩展信息
*/
type TaobaoAilabAicloudTopDeviceExtinfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopDeviceExtinfoGetResponse
}

// 获取设备扩展信息 成功返回结果
type TaobaoAilabAicloudTopDeviceExtinfoGetResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_device_extinfo_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TaobaoAilabAicloudTopDeviceExtinfoGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
