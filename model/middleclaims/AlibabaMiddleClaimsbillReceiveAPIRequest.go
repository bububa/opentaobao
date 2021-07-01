package middleclaims

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMiddleClaimsbillReceiveAPIRequest
国际化中台服务域接收理赔账单 API请求
alibaba.middle.claimsbill.receive

国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔打款的处理后，将该打款结果返回至服务域 */
type AlibabaMiddleClaimsbillReceiveAPIRequest struct {
	model.Params
	// 理赔账单实体
	_claimsBillDto *ClaimsBillDto
}

// New
