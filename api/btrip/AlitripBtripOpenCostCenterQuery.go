package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterQuery 查询成本中心
// alitrip.btrip.open.cost.center.query
//
// 查询成本中心
func AlitripBtripOpenCostCenterQuery(clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterQueryAPIRequest, resp *btrip.AlitripBtripOpenCostCenterQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
