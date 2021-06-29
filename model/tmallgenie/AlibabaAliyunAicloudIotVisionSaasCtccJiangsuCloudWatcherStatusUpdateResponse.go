package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新 APIResponse
alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.cloud.watcher.status.update

天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新
*/
type AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateResponse
}

type AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_aliyun_aicloud_iot_vision_saas_ctcc_jiangsu_cloud_watcher_status_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误码。200 表示成功
    
    RspCode   int64 `json:"rsp_code,omitempty" xml:"rsp_code,omitempty"`

    
    // 一次请求的唯一标识符，和请求中的 seq_id 对齐
    
    SeqId   string `json:"seq_id,omitempty" xml:"seq_id,omitempty"`

    
    // 错误信息。success 表示成功
    
    RspMsg   string `json:"rsp_msg,omitempty" xml:"rsp_msg,omitempty"`

    
}
