package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkPosAfterbuyBenefitSendAPIRequest
生态pos购后发放权益 API请求
alibaba.wdk.pos.afterbuy.benefit.send

生态pos购后发放权益接口开放 */
type AlibabaWdkPosAfterbuyBenefitSendAPIRequest struct {
	model.Params
	// 入参
	_sendBenefitParam *IsvSendBenefitParam
}

// New
