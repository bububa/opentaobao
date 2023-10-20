package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupBigbagWaybillInfo 大包面单查询
// cainiao.global.im.pickup.bigbag.waybill.info
//
// 大包面单查询
func CainiaoGlobalImPickupBigbagWaybillInfo(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest, session string) (*wlbimports.CainiaoGlobalImPickupBigbagWaybillInfoAPIResponse, error) {
	var resp wlbimports.CainiaoGlobalImPickupBigbagWaybillInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
