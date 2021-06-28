package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
openTaoBaoId解绑设备 APIResponse
taobao.ailab.aicloud.top.device.openid.unbind

openTaoBaoId解绑设备
*/
type TaobaoAilabAicloudTopDeviceOpenidUnbindAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ailab_aicloud_top_device_openid_unbind_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"