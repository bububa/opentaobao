package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupBigbagContentCancel 进口大包取消
// cainiao.global.im.pickup.bigbag.content.cancel
//
// 进口大包取消
func CainiaoGlobalImPickupBigbagContentCancel(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupBigbagContentCancelAPIRequest, session string) (*wlbimports.CainiaoGlobalImPickupBigbagContentCancelAPIResponse, error) {
	var resp wlbimports.CainiaoGlobalImPickupBigbagContentCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
