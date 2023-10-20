package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplanlocationquoteupload 9.2-同步地点配额
// alibaba.tmallgenie.scp.plan.location.quote.upload
//
// 同步地点配额
func Alibabatmallgeniescpplanlocationquoteupload(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplanlocationquoteuploadAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplanlocationquoteuploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplanlocationquoteuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
