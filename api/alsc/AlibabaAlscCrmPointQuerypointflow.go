package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmPointQuerypointflow 分页查询积分流水
// alibaba.alsc.crm.point.querypointflow
//
// 分页查询积分流水
func AlibabaAlscCrmPointQuerypointflow(clt *core.SDKClient, req *alsc.AlibabaAlscCrmPointQuerypointflowAPIRequest, resp *alsc.AlibabaAlscCrmPointQuerypointflowAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
