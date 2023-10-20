package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCostCenterTransfer 商旅成本中心转换为外部成本中心
// alitrip.btrip.cost.center.transfer
//
// 商旅成本中心转换为外部成本中心
func AlitripBtripCostCenterTransfer(clt *core.SDKClient, req *btrip.AlitripBtripCostCenterTransferAPIRequest, resp *btrip.AlitripBtripCostCenterTransferAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
