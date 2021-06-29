package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新 API请求
alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.cloud.watcher.status.update

天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新
*/
type AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest struct {
    model.Params
    // 设备唯一标识符
    _ctei   string
    // 设备对应的产品类型
    _devType   string
    // 一次请求的唯一标识符
    _seqId   string
    // 设备所属用户的账号信息
    _userAccount   string
}

// 初始化AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest对象
func NewAlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest() *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest{
    return &AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) GetApiMethodName() string {
    return "alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.cloud.watcher.status.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ctei Setter
// 设备唯一标识符
func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) SetCtei(_ctei string) error {
    r._ctei = _ctei
    r.Set("ctei", _ctei)
    return nil
}

// Ctei Getter
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) GetCtei() string {
    return r._ctei
}
// DevType Setter
// 设备对应的产品类型
func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) SetDevType(_devType string) error {
    r._devType = _devType
    r.Set("dev_type", _devType)
    return nil
}

// DevType Getter
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) GetDevType() string {
    return r._devType
}
// SeqId Setter
// 一次请求的唯一标识符
func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) SetSeqId(_seqId string) error {
    r._seqId = _seqId
    r.Set("seq_id", _seqId)
    return nil
}

// SeqId Getter
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) GetSeqId() string {
    return r._seqId
}
// UserAccount Setter
// 设备所属用户的账号信息
func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) SetUserAccount(_userAccount string) error {
    r._userAccount = _userAccount
    r.Set("user_account", _userAccount)
    return nil
}

// UserAccount Getter
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateRequest) GetUserAccount() string {
    return r._userAccount
}
