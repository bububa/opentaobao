package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripopencostcenterentityadd 增加成本中心人员信息
// alitrip.btrip.open.cost.center.entity.add
//
// 增加成本中心人员信息
func Alitripbtripopencostcenterentityadd(clt *core.SDKClient, req *btrip.AlitripbtripopencostcenterentityaddAPIRequest, session string) (*btrip.AlitripbtripopencostcenterentityaddAPIResponse, error) {
	var resp btrip.AlitripbtripopencostcenterentityaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
