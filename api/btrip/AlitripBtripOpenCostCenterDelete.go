package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

/* AlitripBtripOpenCostCenterDelete
删除成本中心
alitrip.btrip.open.cost.center.delete

删除成本中心 */
func AlitripBtripOpenCostCenterDelete(clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterDeleteAPIRequest, session string) (*btrip.AlitripBtripOpenCostCenterDeleteAPIResponse, error) {
	var resp btrip.AlitripBtripOpenCostCenterDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
