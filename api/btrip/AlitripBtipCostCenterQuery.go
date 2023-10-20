package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtipCostCenterQuery 查询外部成本中心
// alitrip.btip.cost.center.query
//
// 查询外部成本中心
func AlitripBtipCostCenterQuery(clt *core.SDKClient, req *btrip.AlitripBtipCostCenterQueryAPIRequest, resp *btrip.AlitripBtipCostCenterQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
