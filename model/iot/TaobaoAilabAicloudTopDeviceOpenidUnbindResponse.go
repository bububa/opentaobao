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
    Response *TaobaoAilabAicloudTopDeviceOpenidUnbindResponse `json:"taobao_ailab_aicloud_top_device_openid_unbind_response,omitempty"`
}

type TaobaoAilabAicloudTopDeviceOpenidUnbindResponse struct {

    // 结果
    Result   *AiCloudResult `json:"result,omitempty"`

}
