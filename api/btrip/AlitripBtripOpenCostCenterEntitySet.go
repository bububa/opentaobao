package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripopencostcenterentityset 设置成本中心人员信息
// alitrip.btrip.open.cost.center.entity.set
//
// 设置成本中心人员信息
func Alitripbtripopencostcenterentityset(clt *core.SDKClient, req *btrip.AlitripbtripopencostcenterentitysetAPIRequest, session string) (*btrip.AlitripbtripopencostcenterentitysetAPIResponse, error) {
	var resp btrip.AlitripbtripopencostcenterentitysetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
