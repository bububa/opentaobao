package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplannetdemandupload 23-Net Demand净需求回传
// alibaba.tmallgenie.scp.plan.netdemand.upload
//
// Net Demand净需求回传
func Alibabatmallgeniescpplannetdemandupload(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplannetdemanduploadAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplannetdemanduploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplannetdemanduploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
