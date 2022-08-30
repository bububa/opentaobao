package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupBigbagInfo 大包状态查询
// cainiao.global.im.pickup.bigbag.info
//
// 大包状态查询
func CainiaoGlobalImPickupBigbagInfo(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupBigbagInfoAPIRequest, session string) (*wlbimports.CainiaoGlobalImPickupBigbagInfoAPIResponse, error) {
	var resp wlbimports.CainiaoGlobalImPickupBigbagInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
