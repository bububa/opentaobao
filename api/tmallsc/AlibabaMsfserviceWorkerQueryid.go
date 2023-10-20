package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Alibabamsfserviceworkerqueryid 查询师傅workerid
// alibaba.msfservice.worker.queryid
//
// 查询师傅workerid
func Alibabamsfserviceworkerqueryid(clt *core.SDKClient, req *tmallsc.AlibabamsfserviceworkerqueryidAPIRequest, session string) (*tmallsc.AlibabamsfserviceworkerqueryidAPIResponse, error) {
	var resp tmallsc.AlibabamsfserviceworkerqueryidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
