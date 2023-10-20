package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCostCenterDelete 删除外部成本中心
// alitrip.btrip.cost.center.delete
//
// 删除外部成本中心
func AlitripBtripCostCenterDelete(clt *core.SDKClient, req *btrip.AlitripBtripCostCenterDeleteAPIRequest, resp *btrip.AlitripBtripCostCenterDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
