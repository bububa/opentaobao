package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TaobaoCcoSelfCoordinateHandleFinish 天猫国际直购供应商处理完结回复通知
// taobao.cco.self.coordinate.handle.finish
//
// 天猫国际直购供应商处理完结回复通知
func TaobaoCcoSelfCoordinateHandleFinish(clt *core.SDKClient, req *tmallhk.TaobaoCcoSelfCoordinateHandleFinishAPIRequest, session string) (*tmallhk.TaobaoCcoSelfCoordinateHandleFinishAPIResponse, error) {
	var resp tmallhk.TaobaoCcoSelfCoordinateHandleFinishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
