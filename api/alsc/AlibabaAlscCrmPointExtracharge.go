package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmPointExtracharge 积分补录
// alibaba.alsc.crm.point.extracharge
//
// 积分补录
func AlibabaAlscCrmPointExtracharge(clt *core.SDKClient, req *alsc.AlibabaAlscCrmPointExtrachargeAPIRequest, resp *alsc.AlibabaAlscCrmPointExtrachargeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
