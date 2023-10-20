package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkerCreate 服务商工人信息创建
// tmall.servicecenter.worker.create
//
// 服务商工人信息创建
func TmallServicecenterWorkerCreate(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkerCreateAPIRequest, resp *tmallservice.TmallServicecenterWorkerCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
