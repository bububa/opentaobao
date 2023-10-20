package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterModify 修改成本中心
// alitrip.btrip.open.cost.center.modify
//
// 修改成本中心
func AlitripBtripOpenCostCenterModify(clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterModifyAPIRequest, resp *btrip.AlitripBtripOpenCostCenterModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
