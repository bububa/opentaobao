package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaServicecenterWorkcardRelatedskuQuery 查询工单关联的服务项
// alibaba.servicecenter.workcard.relatedsku.query
//
// 查询工单关联的服务项
func AlibabaServicecenterWorkcardRelatedskuQuery(clt *core.SDKClient, req *tmallsc.AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest, resp *tmallsc.AlibabaServicecenterWorkcardRelatedskuQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
