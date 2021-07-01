package middleclaims

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMiddleClaimsacceptReceiveAPIRequest
国际化中台服务域接收保险公司理赔受理结果 API请求
alibaba.middle.claimsaccept.receive

国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔受理结果的处理后，将该结果返回至服务域 */
type AlibabaMiddleClaimsacceptReceiveAPIRequest struct {
	model.Params
	// 理赔受理结果实体类
	_claimsAcceptDto *ClaimsAcceptDto
}

// New
