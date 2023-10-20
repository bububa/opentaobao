package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkerQuery 工人信息查询
// tmall.servicecenter.worker.query
//
// 查询服务商对应的工人信息
func TmallServicecenterWorkerQuery(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkerQueryAPIRequest, session string) (*tmallservice.TmallServicecenterWorkerQueryAPIResponse, error) {
	var resp tmallservice.TmallServicecenterWorkerQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
