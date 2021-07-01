package middleclaims

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMiddleClaimsresultReceiveAPIRequest
国际化中台服务域接收理赔结果 API请求
alibaba.middle.claimsresult.receive

国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔结果的处理后，将该结果返回至服务域 */
type AlibabaMiddleClaimsresultReceiveAPIRequest struct {
	model.Params
	// 理赔结果实体
	_claimsResultDTO *ClaimsResultDto
}

// New
