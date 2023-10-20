package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmPointConsumepoint 积分抵现
// alibaba.alsc.crm.point.consumepoint
//
// 积分抵现
func AlibabaAlscCrmPointConsumepoint(clt *core.SDKClient, req *alsc.AlibabaAlscCrmPointConsumepointAPIRequest, resp *alsc.AlibabaAlscCrmPointConsumepointAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
