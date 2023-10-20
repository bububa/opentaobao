package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupBigbagWaybillInfo 大包面单查询
// cainiao.global.im.pickup.bigbag.waybill.info
//
// 大包面单查询
func CainiaoGlobalImPickupBigbagWaybillInfo(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest, resp *wlbimports.CainiaoGlobalImPickupBigbagWaybillInfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
