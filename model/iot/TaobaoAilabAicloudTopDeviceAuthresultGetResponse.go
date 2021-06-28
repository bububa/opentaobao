package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备授权码验证结果 APIResponse
taobao.ailab.aicloud.top.device.authresult.get

获取设备授权码验证结果
*/
type TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopDeviceAuthresultGetResponse
}

type TaobaoAilabAicloudTopDeviceAuthresultGetResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_device_authresult_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
