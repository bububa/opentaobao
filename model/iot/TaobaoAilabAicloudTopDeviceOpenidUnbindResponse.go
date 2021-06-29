package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
openTaoBaoId解绑设备 API返回值 
taobao.ailab.aicloud.top.device.openid.unbind

openTaoBaoId解绑设备
*/
type TaobaoAilabAicloudTopDeviceOpenidUnbindAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopDeviceOpenidUnbindResponse
}

// openTaoBaoId解绑设备 成功返回结果
type TaobaoAilabAicloudTopDeviceOpenidUnbindResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_device_openid_unbind_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
