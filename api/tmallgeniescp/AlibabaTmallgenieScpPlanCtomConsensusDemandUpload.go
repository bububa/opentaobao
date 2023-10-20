package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplanctomconsensusdemandupload 22-C2M共识需求回传接口
// alibaba.tmallgenie.scp.plan.ctom.consensus.demand.upload
//
// C2M 共识需求回传接口
func Alibabatmallgeniescpplanctomconsensusdemandupload(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplanctomconsensusdemanduploadAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplanctomconsensusdemanduploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplanctomconsensusdemanduploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
