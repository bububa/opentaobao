package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupBigbagContentCreate 大包创建
// cainiao.global.im.pickup.bigbag.content.create
//
// 大包创建
func CainiaoGlobalImPickupBigbagContentCreate(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupBigbagContentCreateAPIRequest, session string) (*wlbimports.CainiaoGlobalImPickupBigbagContentCreateAPIResponse, error) {
	var resp wlbimports.CainiaoGlobalImPickupBigbagContentCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
