package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanChannelGet 5-IBP同步渠道接口
// alibaba.tmallgenie.scp.plan.channel.get
//
// IBP同步渠道接口
func AlibabaTmallgenieScpPlanChannelGet(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanChannelGetAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanChannelGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
