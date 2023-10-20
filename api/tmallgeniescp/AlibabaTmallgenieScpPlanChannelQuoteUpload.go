package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplanchannelquoteupload 9.1-同步渠道配额
// alibaba.tmallgenie.scp.plan.channel.quote.upload
//
// 同步渠道配额
func Alibabatmallgeniescpplanchannelquoteupload(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplanchannelquoteuploadAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplanchannelquoteuploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplanchannelquoteuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
