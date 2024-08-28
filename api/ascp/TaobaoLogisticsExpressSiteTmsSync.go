package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// TaobaoLogisticsExpressSiteTmsSync 配服务商网点信息同步
// taobao.logistics.express.site.tms.sync
//
// 配服务商网点信息同步
func TaobaoLogisticsExpressSiteTmsSync(ctx context.Context, clt *core.SDKClient, req *ascp.TaobaoLogisticsExpressSiteTmsSyncAPIRequest, resp *ascp.TaobaoLogisticsExpressSiteTmsSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
