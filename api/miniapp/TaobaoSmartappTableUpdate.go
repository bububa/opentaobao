package miniapp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoSmartappTableUpdate 智能应用服务登记工作表数据更新
// taobao.smartapp.table.update
//
// 智能应用服务登记工作表数据更新
func TaobaoSmartappTableUpdate(ctx context.Context, clt *core.SDKClient, req *miniapp.TaobaoSmartappTableUpdateAPIRequest, resp *miniapp.TaobaoSmartappTableUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
