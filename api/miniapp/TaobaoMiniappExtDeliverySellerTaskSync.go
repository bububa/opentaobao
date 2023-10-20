package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoMiniappExtDeliverySellerTaskSync 同步外投任务
// taobao.miniapp.ext.delivery.seller.task.sync
//
// 同步外投任务
func TaobaoMiniappExtDeliverySellerTaskSync(clt *core.SDKClient, req *miniapp.TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest, resp *miniapp.TaobaoMiniappExtDeliverySellerTaskSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
