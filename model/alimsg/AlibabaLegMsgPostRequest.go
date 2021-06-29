package alimsg

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
集团法务消息发送 APIRequest
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

func NewAlibabaLegMsgPostRequest() *AlibabaLegMsgPostRequest{
    return &AlibabaLegMsgPostRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegMsgPostRequest) GetApiMethodName() string {
    return "alibaba.leg.msg.post"
}

func (r AlibabaLegMsgPostRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegMsgPostRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r AlibabaLegMsgPostRequest) GetAppId() string {
    return r.appId
}

func (r *AlibabaLegMsgPostRequest) SetAccessKey(accessKey string) error {
    r.accessKey = accessKey
    r.Set("access_key", accessKey)
    return nil
}

func (r AlibabaLegMsgPostRequest) GetAccessKey() string {
    return r.accessKey
}

func (r *AlibabaLegMsgPostRequest) SetMessageDefinitionCode(messageDefinitionCode string) error {
    r.messageDefinitionCode = messageDefinitionCode
    r.Set("message_definition_code", messageDefinitionCode)
    return nil
}

func (r AlibabaLegMsgPostRequest) GetMessageDefinitionCode() string {
    return r.messageDefinitionCode
}

func (r *AlibabaLegMsgPostRequest) SetReceiverType(receiverType string) error {
    r.receiverType = receiverType
    r.Set("receiver_type", receiverType)
    return nil
}

func (r AlibabaLegMsgPostRequest) GetReceiverType() string {
    return r.receiverType
}

func (r *AlibabaLegMsgPostRequest) SetReceivers(receivers string) error {
    r.receivers = receivers
    r.Set("receivers", receivers)
    return nil
}

func (r AlibabaLegMsgPostRequest) GetReceivers() string {
    return r.receivers
}

func (r *AlibabaLegMsgPostRequest) SetMessageBodyListStr(messageBodyListStr string) error {
    r.messageBodyListStr = messageBodyListStr
    r.Set("message_body_list_str", messageBodyListStr)
    return nil
}

func (r AlibabaLegMsgPostRequest) GetMessageBodyListStr() string {
    return r.messageBodyListStr
}

func (r *AlibabaLegMsgPostRequest) SetBusinessId(businessId string) error {
    r.businessId = businessId
    r.Set("business_id", businessId)
    return nil
}

func (r AlibabaLegMsgPostRequest) GetBusinessId() string {
    return r.businessId
}

func (r *AlibabaLegMsgPostRequest) SetBusinessType(businessType string) error {
    r.businessType = businessType
    r.Set("business_type", businessType)
    return nil
}

func (r AlibabaLegMsgPostRequest) GetBusinessType() string {
    return r.businessType
}

func (r *AlibabaLegMsgPostRequest) SetMessageParams(messageParams string) error {
    r.messageParams = messageParams
    r.Set("message_params", messageParams)
    return nil
}

func (r AlibabaLegMsgPostRequest) GetMessageParams() string {
    return r.messageParams
}

func (r *AlibabaLegMsgPostRequest) SetCorpId(corpId string) error {
    r.corpId = corpId
    r.Set("corp_id", corpId)
    return nil
}

func (r AlibabaLegMsgPostRequest) GetCorpId() string {
    return r.corpId
}

func (r *AlibabaLegMsgPostRequest) SetSendTime(sendTime string) error {
    r.sendTime = sendTime
    r.Set("send_time", sendTime)
    return nil
}

func (r AlibabaLegMsgPostRequest) GetSendTime() string {
    return r.sendTime
}

func (r *AlibabaLegMsgPostRequest) SetExpandParamsMapStr(expandParamsMapStr string) error {
    r.expandParamsMapStr = expandParamsMapStr
    r.Set("expand_params_map_str", expandParamsMapStr)
    return nil
}

func (r AlibabaLegMsgPostRequest) GetExpandParamsMapStr() string {
    return r.expandParamsMapStr
}

