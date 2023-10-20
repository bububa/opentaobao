package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// Taobaoccoselfcoordinatebreakorder 天猫国际直购供应商毁单通知
// taobao.cco.self.coordinate.break.order
//
// 天猫国际直购供应商毁单通知
func Taobaoccoselfcoordinatebreakorder(clt *core.SDKClient, req *tmallhk.TaobaoccoselfcoordinatebreakorderAPIRequest, session string) (*tmallhk.TaobaoccoselfcoordinatebreakorderAPIResponse, error) {
	var resp tmallhk.TaobaoccoselfcoordinatebreakorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
