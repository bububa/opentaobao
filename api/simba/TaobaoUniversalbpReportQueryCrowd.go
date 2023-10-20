package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpreportquerycrowd 人群报表查询
// taobao.universalbp.report.query.crowd
//
// 人群报表查询
func Taobaouniversalbpreportquerycrowd(clt *core.SDKClient, req *simba.TaobaouniversalbpreportquerycrowdAPIRequest, session string) (*simba.TaobaouniversalbpreportquerycrowdAPIResponse, error) {
	var resp simba.TaobaouniversalbpreportquerycrowdAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
