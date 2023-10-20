package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardQuery 工单查询接口
// tmall.servicecenter.workcard.query
//
// 工单查询接口
func TmallServicecenterWorkcardQuery(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardQueryAPIRequest, resp *tmallservice.TmallServicecenterWorkcardQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
