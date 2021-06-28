package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取openid设备通用授权码 APIResponse
taobao.ailab.aicloud.top.device.openid.authcode.get

获取openid设备通用授权码
*/
type TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetResponse `json:"ailab_aicloud_top_device_openid_authcode_get_response,omitempty"` 
    TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetResponse struct {

    // 系统自动生成
    
    Result  *struct {
        AiCloudResult  *AiCloudResult `json:"ai_cloud_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetResponse struct {

    // 系统自动生成
    Result   *AiCloudResult `json:"result,omitempty"`

}
