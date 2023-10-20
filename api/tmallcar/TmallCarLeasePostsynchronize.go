package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallcarleasepostsynchronize 天猫开新车租后信息同步
// tmall.car.lease.postsynchronize
//
// 商家同步天猫开新车租后方案
func Tmallcarleasepostsynchronize(clt *core.SDKClient, req *tmallcar.TmallcarleasepostsynchronizeAPIRequest, session string) (*tmallcar.TmallcarleasepostsynchronizeAPIResponse, error) {
	var resp tmallcar.TmallcarleasepostsynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
