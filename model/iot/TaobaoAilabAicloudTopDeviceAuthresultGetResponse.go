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
    Response *TaobaoAilabAicloudTopDeviceAuthresultGetResponse `json:"taobao_ailab_aicloud_top_device_authresult_get_response,omitempty"`
}

type TaobaoAilabAicloudTopDeviceAuthresultGetResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
