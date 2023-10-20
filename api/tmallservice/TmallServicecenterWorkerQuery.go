package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkerQuery 工人信息查询
// tmall.servicecenter.worker.query
//
// 查询服务商对应的工人信息
func TmallServicecenterWorkerQuery(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkerQueryAPIRequest, resp *tmallservice.TmallServicecenterWorkerQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
