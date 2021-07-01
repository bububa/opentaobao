package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaJymRequirementsReceiveAPIRequest
交易猫需求接单接口 API请求
alibaba.jym.requirements.receive

交易猫需求接单接口 */
type AlibabaJymRequirementsReceiveAPIRequest struct {
	model.Params
	// 需求id
	_requirementId string
	// 接单者手机号
	_receiverMobile string
	// 需求订单id
	_requirementOrderId string
}

// New
