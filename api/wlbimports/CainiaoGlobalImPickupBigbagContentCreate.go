package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupBigbagContentCreate 大包创建
// cainiao.global.im.pickup.bigbag.content.create
//
// 大包创建
func CainiaoGlobalImPickupBigbagContentCreate(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupBigbagContentCreateAPIRequest, resp *wlbimports.CainiaoGlobalImPickupBigbagContentCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
