package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// Taobaoccoselfcoordinatehandlefinish 天猫国际直购供应商处理完结回复通知
// taobao.cco.self.coordinate.handle.finish
//
// 天猫国际直购供应商处理完结回复通知
func Taobaoccoselfcoordinatehandlefinish(clt *core.SDKClient, req *tmallhk.TaobaoccoselfcoordinatehandlefinishAPIRequest, session string) (*tmallhk.TaobaoccoselfcoordinatehandlefinishAPIResponse, error) {
	var resp tmallhk.TaobaoccoselfcoordinatehandlefinishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
