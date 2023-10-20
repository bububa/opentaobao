package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterEntitySet 设置成本中心人员信息
// alitrip.btrip.open.cost.center.entity.set
//
// 设置成本中心人员信息
func AlitripBtripOpenCostCenterEntitySet(clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterEntitySetAPIRequest, resp *btrip.AlitripBtripOpenCostCenterEntitySetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
