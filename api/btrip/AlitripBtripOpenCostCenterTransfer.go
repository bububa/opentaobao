package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

/* AlitripBtripOpenCostCenterTransfer
商旅成本中心转换为外部成本中心
alitrip.btrip.open.cost.center.transfer

商旅成本中心转换为外部成本中心 */
func AlitripBtripOpenCostCenterTransfer(clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterTransferAPIRequest, session string) (*btrip.AlitripBtripOpenCostCenterTransferAPIResponse, error) {
	var resp btrip.AlitripBtripOpenCostCenterTransferAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
