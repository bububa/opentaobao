package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCostCenterTransfer 商旅成本中心转换为外部成本中心
// alitrip.btrip.cost.center.transfer
//
// 商旅成本中心转换为外部成本中心
func AlitripBtripCostCenterTransfer(clt *core.SDKClient, req *btrip.AlitripBtripCostCenterTransferAPIRequest, session string) (*btrip.AlitripBtripCostCenterTransferAPIResponse, error) {
	var resp btrip.AlitripBtripCostCenterTransferAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
