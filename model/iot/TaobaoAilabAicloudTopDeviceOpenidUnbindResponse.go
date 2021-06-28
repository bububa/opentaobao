package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
openTaoBaoId解绑设备 APIResponse
taobao.ailab.aicloud.top.device.openid.unbind

openTaoBaoId解绑设备
*/
type TaobaoAilabAicloudTopDeviceOpenidUnbindAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopDeviceOpenidUnbindResponse `json:"ailab_aicloud_top_device_openid_unbind_response,omitempty"` 
    TaobaoAilabAicloudTopDeviceOpenidUnbindResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopDeviceOpenidUnbindResponse struct {

    // 结果
    
    Result  *struct {
        AiCloudResult  *AiCloudResult `json:"ai_cloud_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopDeviceOpenidUnbindResponse struct {

    // 结果
    Result   *AiCloudResult `json:"result,omitempty"`

}
