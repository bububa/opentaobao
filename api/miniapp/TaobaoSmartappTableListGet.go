package miniapp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoSmartappTableListGet 智能应用服务登记工作表列表查询
// taobao.smartapp.table.list.get
//
// 智能应用服务登记工作表列表查询
func TaobaoSmartappTableListGet(ctx context.Context, clt *core.SDKClient, req *miniapp.TaobaoSmartappTableListGetAPIRequest, resp *miniapp.TaobaoSmartappTableListGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
