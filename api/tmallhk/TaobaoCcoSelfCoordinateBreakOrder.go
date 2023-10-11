package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TaobaoCcoSelfCoordinateBreakOrder 天猫国际直购供应商毁单通知
// taobao.cco.self.coordinate.break.order
//
// 天猫国际直购供应商毁单通知
func TaobaoCcoSelfCoordinateBreakOrder(clt *core.SDKClient, req *tmallhk.TaobaoCcoSelfCoordinateBreakOrderAPIRequest, session string) (*tmallhk.TaobaoCcoSelfCoordinateBreakOrderAPIResponse, error) {
	var resp tmallhk.TaobaoCcoSelfCoordinateBreakOrderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
