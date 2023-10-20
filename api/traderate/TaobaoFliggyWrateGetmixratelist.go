package traderate

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traderate"
)

// Taobaofliggywrategetmixratelist 飞猪通用评价接口
// taobao.fliggy.wrate.getmixratelist
//
// 飞猪评价通用接口
func Taobaofliggywrategetmixratelist(clt *core.SDKClient, req *traderate.TaobaofliggywrategetmixratelistAPIRequest, session string) (*traderate.TaobaofliggywrategetmixratelistAPIResponse, error) {
	var resp traderate.TaobaofliggywrategetmixratelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
