package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取设备授权码验证结果 APIResponse
taobao.ailab.aicloud.top.device.authresult.get

获取设备授权码验证结果
*/
type TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopDeviceAuthresultGetResponse `json:"ailab_aicloud_top_device_authresult_get_response,omitempty"` 
    TaobaoAilabAicloudTopDeviceAuthresultGetResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopDeviceAuthresultGetResponse struct {

    // result
    
    Result  *struct {
        AiCloudResult  *AiCloudResult `json:"ai_cloud_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopDeviceAuthresultGetResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
