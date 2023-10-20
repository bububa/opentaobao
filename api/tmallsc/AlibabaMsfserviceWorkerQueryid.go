package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaMsfserviceWorkerQueryid 查询师傅workerid
// alibaba.msfservice.worker.queryid
//
// 查询师傅workerid
func AlibabaMsfserviceWorkerQueryid(clt *core.SDKClient, req *tmallsc.AlibabaMsfserviceWorkerQueryidAPIRequest, session string) (*tmallsc.AlibabaMsfserviceWorkerQueryidAPIResponse, error) {
	var resp tmallsc.AlibabaMsfserviceWorkerQueryidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
