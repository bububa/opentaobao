package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallcarleasecitysynchronize 天猫开新车租后分期城市信息同步
// tmall.car.lease.citysynchronize
//
// 天猫开新车租后分期城市信息同步
func Tmallcarleasecitysynchronize(clt *core.SDKClient, req *tmallcar.TmallcarleasecitysynchronizeAPIRequest, session string) (*tmallcar.TmallcarleasecitysynchronizeAPIResponse, error) {
	var resp tmallcar.TmallcarleasecitysynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
