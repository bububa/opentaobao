package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallcarleasestatussynchronize 天猫开新车租后状态同步
// tmall.car.lease.statussynchronize
//
// 天猫开新车租后状态同步
func Tmallcarleasestatussynchronize(clt *core.SDKClient, req *tmallcar.TmallcarleasestatussynchronizeAPIRequest, session string) (*tmallcar.TmallcarleasestatussynchronizeAPIResponse, error) {
	var resp tmallcar.TmallcarleasestatussynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
