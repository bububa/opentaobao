package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取openId设备授权码验证结果 APIResponse
taobao.ailab.aicloud.top.device.openid.authresult.get

获取openId设备授权码验证结果
*/
type TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetResponse `json:"ailab_aicloud_top_device_openid_authresult_get_response,omitempty"` 
    TaobaoAilabAicloudTopDeviceOpenidAuthresultGetResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopDeviceOpenidAuthresultGetResponse struct {

    // 系统自动生成
    
    Result  *struct {
        AiCloudResult  *AiCloudResult `json:"ai_cloud_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopDeviceOpenidAuthresultGetResponse struct {

    // 系统自动生成
    Result   *AiCloudResult `json:"result,omitempty"`

}
