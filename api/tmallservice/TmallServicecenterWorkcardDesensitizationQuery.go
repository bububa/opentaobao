package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardDesensitizationQuery 工单查询接口
// tmall.servicecenter.workcard.desensitization.query
//
// 工单查询接口
func TmallServicecenterWorkcardDesensitizationQuery(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardDesensitizationQueryAPIRequest, resp *tmallservice.TmallServicecenterWorkcardDesensitizationQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
