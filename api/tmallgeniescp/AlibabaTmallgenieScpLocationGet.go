package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpLocationGet 2-IBP查询CDC和RDC数据接口
// alibaba.tmallgenie.scp.location.get
//
// 天猫精灵供应链-计划域-IBP查询CDC和RDC数据接口
func AlibabaTmallgenieScpLocationGet(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpLocationGetAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpLocationGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
