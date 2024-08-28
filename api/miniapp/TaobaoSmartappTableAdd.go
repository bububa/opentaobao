package miniapp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoSmartappTableAdd 智能应用服务登记工作表数据新增
// taobao.smartapp.table.add
//
// 智能应用服务登记工作表数据新增
func TaobaoSmartappTableAdd(ctx context.Context, clt *core.SDKClient, req *miniapp.TaobaoSmartappTableAddAPIRequest, resp *miniapp.TaobaoSmartappTableAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
