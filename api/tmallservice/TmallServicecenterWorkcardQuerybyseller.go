package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardQuerybyseller 工单查询接口（面向商家）
// tmall.servicecenter.workcard.querybyseller
//
// 查询工单
func TmallServicecenterWorkcardQuerybyseller(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardQuerybysellerAPIRequest, resp *tmallservice.TmallServicecenterWorkcardQuerybysellerAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
