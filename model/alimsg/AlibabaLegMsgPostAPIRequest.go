package alimsg

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegmsgpostAPIRequest 集团法务消息发送 API请求
// alibaba.leg.msg.post
//
// 消息发送能力
type AlibabalegmsgpostAPIRequest struct {
	model.Params
	// 应用标识
	_appId string
	// 认证的code
	_accessKey string
	// 消息定义code
	_messageDefinitionCode string
	// 接收人类型
	_receiverType string
	// 接收人数组
	_receivers string
	// 发送的渠道类型数组
	_messageBodyListStr string
	// 业务id
	_businessId string
	// 业务类型
	_businessType string
	// 模版里定义的变量
	_messageParams string
	// 三方租户id
	_corpId string
	// 发送时间
	_sendTime string
	// 扩展参数
	_expandParamsMapStr string
}

// NewAlibabalegmsgpostRequest 初始化AlibabalegmsgpostAPIRequest对象
func NewAlibabalegmsgpostRequest() *AlibabalegmsgpostAPIRequest {
	return &AlibabalegmsgpostAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegmsgpostAPIRequest) GetApiMethodName() string {
	return "alibaba.leg.msg.post"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegmsgpostAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegmsgpostAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppId is AppId Setter
// 应用标识
func (r *AlibabalegmsgpostAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabalegmsgpostAPIRequest) GetAppId() string {
	return r._appId
}

// SetAccessKey is AccessKey Setter
// 认证的code
func (r *AlibabalegmsgpostAPIRequest) SetAccessKey(_accessKey string) error {
	r._accessKey = _accessKey
	r.Set("access_key", _accessKey)
	return nil
}

// GetAccessKey AccessKey Getter
func (r AlibabalegmsgpostAPIRequest) GetAccessKey() string {
	return r._accessKey
}

// SetMessageDefinitionCode is MessageDefinitionCode Setter
// 消息定义code
func (r *AlibabalegmsgpostAPIRequest) SetMessageDefinitionCode(_messageDefinitionCode string) error {
	r._messageDefinitionCode = _messageDefinitionCode
	r.Set("message_definition_code", _messageDefinitionCode)
	return nil
}

// GetMessageDefinitionCode MessageDefinitionCode Getter
func (r AlibabalegmsgpostAPIRequest) GetMessageDefinitionCode() string {
	return r._messageDefinitionCode
}

// SetReceiverType is ReceiverType Setter
// 接收人类型
func (r *AlibabalegmsgpostAPIRequest) SetReceiverType(_receiverType string) error {
	r._receiverType = _receiverType
	r.Set("receiver_type", _receiverType)
	return nil
}

// GetReceiverType ReceiverType Getter
func (r AlibabalegmsgpostAPIRequest) GetReceiverType() string {
	return r._receiverType
}

// SetReceivers is Receivers Setter
// 接收人数组
func (r *AlibabalegmsgpostAPIRequest) SetReceivers(_receivers string) error {
	r._receivers = _receivers
	r.Set("receivers", _receivers)
	return nil
}

// GetReceivers Receivers Getter
func (r AlibabalegmsgpostAPIRequest) GetReceivers() string {
	return r._receivers
}

// SetMessageBodyListStr is MessageBodyListStr Setter
// 发送的渠道类型数组
func (r *AlibabalegmsgpostAPIRequest) SetMessageBodyListStr(_messageBodyListStr string) error {
	r._messageBodyListStr = _messageBodyListStr
	r.Set("message_body_list_str", _messageBodyListStr)
	return nil
}

// GetMessageBodyListStr MessageBodyListStr Getter
func (r AlibabalegmsgpostAPIRequest) GetMessageBodyListStr() string {
	return r._messageBodyListStr
}

// SetBusinessId is BusinessId Setter
// 业务id
func (r *AlibabalegmsgpostAPIRequest) SetBusinessId(_businessId string) error {
	r._businessId = _businessId
	r.Set("business_id", _businessId)
	return nil
}

// GetBusinessId BusinessId Getter
func (r AlibabalegmsgpostAPIRequest) GetBusinessId() string {
	return r._businessId
}

// SetBusinessType is BusinessType Setter
// 业务类型
func (r *AlibabalegmsgpostAPIRequest) SetBusinessType(_businessType string) error {
	r._businessType = _businessType
	r.Set("business_type", _businessType)
	return nil
}

// GetBusinessType BusinessType Getter
func (r AlibabalegmsgpostAPIRequest) GetBusinessType() string {
	return r._businessType
}

// SetMessageParams is MessageParams Setter
// 模版里定义的变量
func (r *AlibabalegmsgpostAPIRequest) SetMessageParams(_messageParams string) error {
	r._messageParams = _messageParams
	r.Set("message_params", _messageParams)
	return nil
}

// GetMessageParams MessageParams Getter
func (r AlibabalegmsgpostAPIRequest) GetMessageParams() string {
	return r._messageParams
}

// SetCorpId is CorpId Setter
// 三方租户id
func (r *AlibabalegmsgpostAPIRequest) SetCorpId(_corpId string) error {
	r._corpId = _corpId
	r.Set("corp_id", _corpId)
	return nil
}

// GetCorpId CorpId Getter
func (r AlibabalegmsgpostAPIRequest) GetCorpId() string {
	return r._corpId
}

// SetSendTime is SendTime Setter
// 发送时间
func (r *AlibabalegmsgpostAPIRequest) SetSendTime(_sendTime string) error {
	r._sendTime = _sendTime
	r.Set("send_time", _sendTime)
	return nil
}

// GetSendTime SendTime Getter
func (r AlibabalegmsgpostAPIRequest) GetSendTime() string {
	return r._sendTime
}

// SetExpandParamsMapStr is ExpandParamsMapStr Setter
// 扩展参数
func (r *AlibabalegmsgpostAPIRequest) SetExpandParamsMapStr(_expandParamsMapStr string) error {
	r._expandParamsMapStr = _expandParamsMapStr
	r.Set("expand_params_map_str", _expandParamsMapStr)
	return nil
}

// GetExpandParamsMapStr ExpandParamsMapStr Getter
func (r AlibabalegmsgpostAPIRequest) GetExpandParamsMapStr() string {
	return r._expandParamsMapStr
}
