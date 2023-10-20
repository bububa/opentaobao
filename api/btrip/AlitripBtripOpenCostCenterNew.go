package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterNew 新增成本中心
// alitrip.btrip.open.cost.center.new
//
// 新增成本中心
func AlitripBtripOpenCostCenterNew(clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterNewAPIRequest, resp *btrip.AlitripBtripOpenCostCenterNewAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
