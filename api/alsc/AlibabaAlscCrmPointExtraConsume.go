package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmPointExtraConsume 积分补扣
// alibaba.alsc.crm.point.extra.consume
//
// 积分补扣
func AlibabaAlscCrmPointExtraConsume(clt *core.SDKClient, req *alsc.AlibabaAlscCrmPointExtraConsumeAPIRequest, resp *alsc.AlibabaAlscCrmPointExtraConsumeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
