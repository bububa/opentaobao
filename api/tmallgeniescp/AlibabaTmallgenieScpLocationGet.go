package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpLocationGet 2-IBP查询CDC和RDC数据接口
// alibaba.tmallgenie.scp.location.get
//
// 天猫精灵供应链-计划域-IBP查询CDC和RDC数据接口
func AlibabaTmallgenieScpLocationGet(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpLocationGetAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpLocationGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
