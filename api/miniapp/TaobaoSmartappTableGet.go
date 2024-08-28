package miniapp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoSmartappTableGet 智能应用服务登记工作表数据查询
// taobao.smartapp.table.get
//
// 智能应用服务登记工作表数据查询
func TaobaoSmartappTableGet(ctx context.Context, clt *core.SDKClient, req *miniapp.TaobaoSmartappTableGetAPIRequest, resp *miniapp.TaobaoSmartappTableGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
