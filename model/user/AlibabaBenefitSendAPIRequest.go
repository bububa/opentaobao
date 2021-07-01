package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaBenefitSendAPIRequest
发奖接口 API请求
alibaba.benefit.send

发奖接口 */
type AlibabaBenefitSendAPIRequest struct {
	model.Params
	// 发放的权益(奖品)唯一标识
	_rightEname string
	// 接收奖品的用户openId
	_receiverId string
	// 规定填taobao即可
	_userType string
	// 幂等校验id，业务重试需要，自定义唯一字段即可
	_uniqueId string
	// mtop
	_appName string
	// 调用应用ip，非必填
	_ip string
}

// New
