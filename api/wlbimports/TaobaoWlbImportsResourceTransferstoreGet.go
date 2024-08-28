package wlbimports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// TaobaoWlbImportsResourceTransferstoreGet 根据指定的资源获取所有中转仓列表
// taobao.wlb.imports.resource.transferstore.get
//
// 根据指定的资源获取所有中转仓列表
func TaobaoWlbImportsResourceTransferstoreGet(ctx context.Context, clt *core.SDKClient, req *wlbimports.TaobaoWlbImportsResourceTransferstoreGetAPIRequest, resp *wlbimports.TaobaoWlbImportsResourceTransferstoreGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
