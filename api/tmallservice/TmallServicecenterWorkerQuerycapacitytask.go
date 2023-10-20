package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkerQuerycapacitytask 查询需求容量
// tmall.servicecenter.worker.querycapacitytask
//
// 查询需求容量
func TmallServicecenterWorkerQuerycapacitytask(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkerQuerycapacitytaskAPIRequest, resp *tmallservice.TmallServicecenterWorkerQuerycapacitytaskAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
