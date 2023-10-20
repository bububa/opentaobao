package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterSpserviceorderQuery 服务单列表查询
// tmall.servicecenter.spserviceorder.query
//
// 查询服务单列表
func TmallServicecenterSpserviceorderQuery(clt *core.SDKClient, req *tmallservice.TmallServicecenterSpserviceorderQueryAPIRequest, resp *tmallservice.TmallServicecenterSpserviceorderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
