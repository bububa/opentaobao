package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaServicecenterWorkcardConfirmedskuQuery 查询确认履行的服务项
// alibaba.servicecenter.workcard.confirmedsku.query
//
// 查询确认履行的服务项
func AlibabaServicecenterWorkcardConfirmedskuQuery(clt *core.SDKClient, req *tmallsc.AlibabaServicecenterWorkcardConfirmedskuQueryAPIRequest, resp *tmallsc.AlibabaServicecenterWorkcardConfirmedskuQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
