package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新 APIRequest
alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.cloud.watcher.status.update

天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新
*/
type AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest struct {
    model.Params

    // 设备唯一标识符
    ctei   string 

    // 设备对应的产品类型
    devType   string 

    // 一次请求的唯一标识符
    seqId   string 

    // 设备所属用户的账号信息
    userAccount   string 

}

func NewAlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest() *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest{
    return &AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) GetApiMethodName() string {
    return "alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.cloud.watcher.status.update"
}

func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) SetCtei(ctei string) error {
    r.ctei = ctei
    r.Set("ctei", ctei)
    return nil
}

func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) GetCtei() string {
    return r.ctei
}

func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) SetDevType(devType string) error {
    r.devType = devType
    r.Set("dev_type", devType)
    return nil
}

func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) GetDevType() string {
    return r.devType
}

func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) SetSeqId(seqId string) error {
    r.seqId = seqId
    r.Set("seq_id", seqId)
    return nil
}

func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) GetSeqId() string {
    return r.seqId
}

func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) SetUserAccount(userAccount string) error {
    r.userAccount = userAccount
    r.Set("user_account", userAccount)
    return nil
}

func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) GetUserAccount() string {
    return r.userAccount
}

