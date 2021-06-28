package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取openId设备授权码验证结果 APIResponse
taobao.ailab.aicloud.top.device.openid.authresult.get

获取openId设备授权码验证结果
*/
type TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopDeviceOpenidAuthresultGetResponse
}

type TaobaoAilabAicloudTopDeviceOpenidAuthresultGetResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_device_openid_authresult_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 系统自动生成
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
