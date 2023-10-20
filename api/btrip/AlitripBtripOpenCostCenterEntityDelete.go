package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripopencostcenterentitydelete 删除成本中心人员信息
// alitrip.btrip.open.cost.center.entity.delete
//
// 删除成本中心人员信息
func Alitripbtripopencostcenterentitydelete(clt *core.SDKClient, req *btrip.AlitripbtripopencostcenterentitydeleteAPIRequest, session string) (*btrip.AlitripbtripopencostcenterentitydeleteAPIResponse, error) {
	var resp btrip.AlitripbtripopencostcenterentitydeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
