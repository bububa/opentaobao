package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkerQuerycapacitytask 查询需求容量
// tmall.servicecenter.worker.querycapacitytask
//
// 查询需求容量
func TmallServicecenterWorkerQuerycapacitytask(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkerQuerycapacitytaskAPIRequest, session string) (*tmallservice.TmallServicecenterWorkerQuerycapacitytaskAPIResponse, error) {
	var resp tmallservice.TmallServicecenterWorkerQuerycapacitytaskAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
