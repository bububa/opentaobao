package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupStoresGet 首公里揽收-集货仓列表查询
// cainiao.global.im.pickup.stores.get
//
// 首公里揽收-集货仓列表查询
func CainiaoGlobalImPickupStoresGet(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupStoresGetAPIRequest, session string) (*wlbimports.CainiaoGlobalImPickupStoresGetAPIResponse, error) {
	var resp wlbimports.CainiaoGlobalImPickupStoresGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
