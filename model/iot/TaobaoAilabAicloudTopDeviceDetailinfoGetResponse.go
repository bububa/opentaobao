package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备详细信息 APIResponse
taobao.ailab.aicloud.top.device.detailinfo.get

获取设备详细信息
*/
type TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopDeviceDetailinfoGetResponse
}

type TaobaoAilabAicloudTopDeviceDetailinfoGetResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_device_detailinfo_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *TaobaoAilabAicloudTopDeviceDetailinfoGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
