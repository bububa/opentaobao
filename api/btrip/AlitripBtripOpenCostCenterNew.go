package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterNew 新增成本中心
// alitrip.btrip.open.cost.center.new
//
// 新增成本中心
func AlitripBtripOpenCostCenterNew(clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterNewAPIRequest, session string) (*btrip.AlitripBtripOpenCostCenterNewAPIResponse, error) {
	var resp btrip.AlitripBtripOpenCostCenterNewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
