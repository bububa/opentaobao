package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备状态信息 APIResponse
taobao.ailab.aicloud.top.device.statusinfo.get

获取设备状态信息
*/
type TaobaoAilabAicloudTopDeviceStatusinfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopDeviceStatusinfoGetResponse
}

type TaobaoAilabAicloudTopDeviceStatusinfoGetResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_device_statusinfo_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoAilabAicloudTopDeviceStatusinfoGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
