package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterEntityAdd 增加成本中心人员信息
// alitrip.btrip.open.cost.center.entity.add
//
// 增加成本中心人员信息
func AlitripBtripOpenCostCenterEntityAdd(clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterEntityAddAPIRequest, resp *btrip.AlitripBtripOpenCostCenterEntityAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
