package alimsg

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
集团法务消息发送 API请求
alibaba.leg.msg.post

消息发送能力
*/
type AlibabaLegMsgPostRequest struct {
    model.Params
    // 应用标识
    appId   string
    // 认证的code
    accessKey   string
    // 消息定义code
    messageDefinitionCode   string
    // 接收人类型
    receiverType   string
    // 接收人数组
    receivers   string
    // 发送的渠道类型数组
    messageBodyListStr   string
    // 业务id
    businessId   string
    // 业务类型
    businessType   string
    // 模版里定义的变量
    messageParams   string
    // 三方租户id
    corpId   string
    // 发送时间
    sendTime   string
    // 扩展参数
    expandParamsMapStr   string
}

// 初始化AlibabaLegMsgPostRequest对象
func NewAlibabaLegMsgPostRequest() *AlibabaLegMsgPostRequest{
    return &AlibabaLegMsgPostRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegMsgPostRequest) GetApiMethodName() string {
    return "alibaba.leg.msg.post"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegMsgPostRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppId Setter
// 应用标识
func (r *AlibabaLegMsgPostRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

// AppId Getter
func (r AlibabaLegMsgPostRequest) GetAppId() string {
    return r.appId
}
// AccessKey Setter
// 认证的code
func (r *AlibabaLegMsgPostRequest) SetAccessKey(accessKey string) error {
    r.accessKey = accessKey
    r.Set("access_key", accessKey)
    return nil
}

// AccessKey Getter
func (r AlibabaLegMsgPostRequest) GetAccessKey() string {
    return r.accessKey
}
// MessageDefinitionCode Setter
// 消息定义code
func (r *AlibabaLegMsgPostRequest) SetMessageDefinitionCode(messageDefinitionCode string) error {
    r.messageDefinitionCode = messageDefinitionCode
    r.Set("message_definition_code", messageDefinitionCode)
    return nil
}

// MessageDefinitionCode Getter
func (r AlibabaLegMsgPostRequest) GetMessageDefinitionCode() string {
    return r.messageDefinitionCode
}
// ReceiverType Setter
// 接收人类型
func (r *AlibabaLegMsgPostRequest) SetReceiverType(receiverType string) error {
    r.receiverType = receiverType
    r.Set("receiver_type", receiverType)
    return nil
}

// ReceiverType Getter
func (r AlibabaLegMsgPostRequest) GetReceiverType() string {
    return r.receiverType
}
// Receivers Setter
// 接收人数组
func (r *AlibabaLegMsgPostRequest) SetReceivers(receivers string) error {
    r.receivers = receivers
    r.Set("receivers", receivers)
    return nil
}

// Receivers Getter
func (r AlibabaLegMsgPostRequest) GetReceivers() string {
    return r.receivers
}
// MessageBodyListStr Setter
// 发送的渠道类型数组
func (r *AlibabaLegMsgPostRequest) SetMessageBodyListStr(messageBodyListStr string) error {
    r.messageBodyListStr = messageBodyListStr
    r.Set("message_body_list_str", messageBodyListStr)
    return nil
}

// MessageBodyListStr Getter
func (r AlibabaLegMsgPostRequest) GetMessageBodyListStr() string {
    return r.messageBodyListStr
}
// BusinessId Setter
// 业务id
func (r *AlibabaLegMsgPostRequest) SetBusinessId(businessId string) error {
    r.businessId = businessId
    r.Set("business_id", businessId)
    return nil
}

// BusinessId Getter
func (r AlibabaLegMsgPostRequest) GetBusinessId() string {
    return r.businessId
}
// BusinessType Setter
// 业务类型
func (r *AlibabaLegMsgPostRequest) SetBusinessType(businessType string) error {
    r.businessType = businessType
    r.Set("business_type", businessType)
    return nil
}

// BusinessType Getter
func (r AlibabaLegMsgPostRequest) GetBusinessType() string {
    return r.businessType
}
// MessageParams Setter
// 模版里定义的变量
func (r *AlibabaLegMsgPostRequest) SetMessageParams(messageParams string) error {
    r.messageParams = messageParams
    r.Set("message_params", messageParams)
    return nil
}

// MessageParams Getter
func (r AlibabaLegMsgPostRequest) GetMessageParams() string {
    return r.messageParams
}
// CorpId Setter
// 三方租户id
func (r *AlibabaLegMsgPostRequest) SetCorpId(corpId string) error {
    r.corpId = corpId
    r.Set("corp_id", corpId)
    return nil
}

// CorpId Getter
func (r AlibabaLegMsgPostRequest) GetCorpId() string {
    return r.corpId
}
// SendTime Setter
// 发送时间
func (r *AlibabaLegMsgPostRequest) SetSendTime(sendTime string) error {
    r.sendTime = sendTime
    r.Set("send_time", sendTime)
    return nil
}

// SendTime Getter
func (r AlibabaLegMsgPostRequest) GetSendTime() string {
    return r.sendTime
}
// ExpandParamsMapStr Setter
// 扩展参数
func (r *AlibabaLegMsgPostRequest) SetExpandParamsMapStr(expandParamsMapStr string) error {
    r.expandParamsMapStr = expandParamsMapStr
    r.Set("expand_params_map_str", expandParamsMapStr)
    return nil
}

// ExpandParamsMapStr Getter
func (r AlibabaLegMsgPostRequest) GetExpandParamsMapStr() string {
    return r.expandParamsMapStr
}
