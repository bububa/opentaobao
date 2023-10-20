package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupBigbagContentCancel 进口大包取消
// cainiao.global.im.pickup.bigbag.content.cancel
//
// 进口大包取消
func CainiaoGlobalImPickupBigbagContentCancel(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupBigbagContentCancelAPIRequest, resp *wlbimports.CainiaoGlobalImPickupBigbagContentCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
