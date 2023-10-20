package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanChannelQuoteUpload 9.1-同步渠道配额
// alibaba.tmallgenie.scp.plan.channel.quote.upload
//
// 同步渠道配额
func AlibabaTmallgenieScpPlanChannelQuoteUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanChannelQuoteUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
