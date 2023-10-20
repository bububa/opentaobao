package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterSpserviceorderCreate 服务单创建
// tmall.servicecenter.spserviceorder.create
//
// 服务单创建
func TmallServicecenterSpserviceorderCreate(clt *core.SDKClient, req *tmallservice.TmallServicecenterSpserviceorderCreateAPIRequest, resp *tmallservice.TmallServicecenterSpserviceorderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
