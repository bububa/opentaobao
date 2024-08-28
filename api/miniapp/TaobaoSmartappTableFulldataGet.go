package miniapp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoSmartappTableFulldataGet 智能应用工作表地址查询
// taobao.smartapp.table.fulldata.get
//
// 智能应用工作表地址查询
func TaobaoSmartappTableFulldataGet(ctx context.Context, clt *core.SDKClient, req *miniapp.TaobaoSmartappTableFulldataGetAPIRequest, resp *miniapp.TaobaoSmartappTableFulldataGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
