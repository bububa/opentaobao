package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵 IoT 视频 SaaS 服务-江苏电信-appKeySecret 更新 API请求
alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.key.secret.update

天猫精灵 IoT 视频 SaaS 服务-江苏电信-appKeySecret 更新
*/
type AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest struct {
    model.Params
    // 一次请求的唯一标识符
    _seqId   string
    // 新的 key
    _secret   string
}

// 初始化AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest对象
func NewAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateRequest() *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest{
    return &AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.key.secret.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SeqId Setter
// 一次请求的唯一标识符
func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest) SetSeqId(_seqId string) error {
    r._seqId = _seqId
    r.Set("seq_id", _seqId)
    return nil
}

// SeqId Getter
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest) GetSeqId() string {
    return r._seqId
}
// Secret Setter
// 新的 key
func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest) SetSecret(_secret string) error {
    r._secret = _secret
    r.Set("secret", _secret)
    return nil
}

// Secret Getter
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest) GetSecret() string {
    return r._secret
}
