package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplanconsensusdemandupload 20-IBP共识需求回传接口
// alibaba.tmallgenie.scp.plan.consensus.demand.upload
//
// IBP共识需求回传接口
func Alibabatmallgeniescpplanconsensusdemandupload(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplanconsensusdemanduploadAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplanconsensusdemanduploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplanconsensusdemanduploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
