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
    _appId   string
    // 认证的code
    _accessKey   string
    // 消息定义code
    _messageDefinitionCode   string
    // 接收人类型
    _receiverType   string
    // 接收人数组
    _receivers   string
    // 发送的渠道类型数组
    _messageBodyListStr   string
    // 业务id
    _businessId   string
    // 业务类型
    _businessType   string
    // 模版里定义的变量
    _messageParams   string
    // 三方租户id
    _corpId   string
    // 发送时间
    _sendTime   string
    // 扩展参数
    _expandParamsMapStr   string
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
func (r *AlibabaLegMsgPostRequest) SetAppId(_appId string) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r AlibabaLegMsgPostRequest) GetAppId() string {
    return r._appId
}
// AccessKey Setter
// 认证的code
func (r *AlibabaLegMsgPostRequest) SetAccessKey(_accessKey string) error {
    r._accessKey = _accessKey
    r.Set("access_key", _accessKey)
    return nil
}

// AccessKey Getter
func (r AlibabaLegMsgPostRequest) GetAccessKey() string {
    return r._accessKey
}
// MessageDefinitionCode Setter
// 消息定义code
func (r *AlibabaLegMsgPostRequest) SetMessageDefinitionCode(_messageDefinitionCode string) error {
    r._messageDefinitionCode = _messageDefinitionCode
    r.Set("message_definition_code", _messageDefinitionCode)
    return nil
}

// MessageDefinitionCode Getter
func (r AlibabaLegMsgPostRequest) GetMessageDefinitionCode() string {
    return r._messageDefinitionCode
}
// ReceiverType Setter
// 接收人类型
func (r *AlibabaLegMsgPostRequest) SetReceiverType(_receiverType string) error {
    r._receiverType = _receiverType
    r.Set("receiver_type", _receiverType)
    return nil
}

// ReceiverType Getter
func (r AlibabaLegMsgPostRequest) GetReceiverType() string {
    return r._receiverType
}
// Receivers Setter
// 接收人数组
func (r *AlibabaLegMsgPostRequest) SetReceivers(_receivers string) error {
    r._receivers = _receivers
    r.Set("receivers", _receivers)
    return nil
}

// Receivers Getter
func (r AlibabaLegMsgPostRequest) GetReceivers() string {
    return r._receivers
}
// MessageBodyListStr Setter
// 发送的渠道类型数组
func (r *AlibabaLegMsgPostRequest) SetMessageBodyListStr(_messageBodyListStr string) error {
    r._messageBodyListStr = _messageBodyListStr
    r.Set("message_body_list_str", _messageBodyListStr)
    return nil
}

// MessageBodyListStr Getter
func (r AlibabaLegMsgPostRequest) GetMessageBodyListStr() string {
    return r._messageBodyListStr
}
// BusinessId Setter
// 业务id
func (r *AlibabaLegMsgPostRequest) SetBusinessId(_businessId string) error {
    r._businessId = _businessId
    r.Set("business_id", _businessId)
    return nil
}

// BusinessId Getter
func (r AlibabaLegMsgPostRequest) GetBusinessId() string {
    return r._businessId
}
// BusinessType Setter
// 业务类型
func (r *AlibabaLegMsgPostRequest) SetBusinessType(_businessType string) error {
    r._businessType = _businessType
    r.Set("business_type", _businessType)
    return nil
}

// BusinessType Getter
func (r AlibabaLegMsgPostRequest) GetBusinessType() string {
    return r._businessType
}
// MessageParams Setter
// 模版里定义的变量
func (r *AlibabaLegMsgPostRequest) SetMessageParams(_messageParams string) error {
    r._messageParams = _messageParams
    r.Set("message_params", _messageParams)
    return nil
}

// MessageParams Getter
func (r AlibabaLegMsgPostRequest) GetMessageParams() string {
    return r._messageParams
}
// CorpId Setter
// 三方租户id
func (r *AlibabaLegMsgPostRequest) SetCorpId(_corpId string) error {
    r._corpId = _corpId
    r.Set("corp_id", _corpId)
    return nil
}

// CorpId Getter
func (r AlibabaLegMsgPostRequest) GetCorpId() string {
    return r._corpId
}
// SendTime Setter
// 发送时间
func (r *AlibabaLegMsgPostRequest) SetSendTime(_sendTime string) error {
    r._sendTime = _sendTime
    r.Set("send_time", _sendTime)
    return nil
}

// SendTime Getter
func (r AlibabaLegMsgPostRequest) GetSendTime() string {
    return r._sendTime
}
// ExpandParamsMapStr Setter
// 扩展参数
func (r *AlibabaLegMsgPostRequest) SetExpandParamsMapStr(_expandParamsMapStr string) error {
    r._expandParamsMapStr = _expandParamsMapStr
    r.Set("expand_params_map_str", _expandParamsMapStr)
    return nil
}

// ExpandParamsMapStr Getter
func (r AlibabaLegMsgPostRequest) GetExpandParamsMapStr() string {
    return r._expandParamsMapStr
}
