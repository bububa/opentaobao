package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterTransfer 商旅成本中心转换为外部成本中心
// alitrip.btrip.open.cost.center.transfer
//
// 商旅成本中心转换为外部成本中心
func AlitripBtripOpenCostCenterTransfer(clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterTransferAPIRequest, resp *btrip.AlitripBtripOpenCostCenterTransferAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
