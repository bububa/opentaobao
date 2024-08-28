package wlb

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbWmsInventoryLackUpload 缺货通知
// taobao.wlb.wms.inventory.lack.upload
//
// 缺货通知
func TaobaoWlbWmsInventoryLackUpload(ctx context.Context, clt *core.SDKClient, req *wlb.TaobaoWlbWmsInventoryLackUploadAPIRequest, resp *wlb.TaobaoWlbWmsInventoryLackUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
