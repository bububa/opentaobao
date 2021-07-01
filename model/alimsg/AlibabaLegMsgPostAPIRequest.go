package alimsg

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegMsgPostAPIRequest
集团法务消息发送 API请求
alibaba.leg.msg.post

消息发送能力 */
type AlibabaLegMsgPostAPIRequest struct {
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

// New
