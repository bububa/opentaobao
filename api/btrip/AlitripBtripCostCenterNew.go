package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCostCenterNew 新建外部成本中心
// alitrip.btrip.cost.center.new
//
// 新建外部成本中心
func AlitripBtripCostCenterNew(clt *core.SDKClient, req *btrip.AlitripBtripCostCenterNewAPIRequest, session string) (*btrip.AlitripBtripCostCenterNewAPIResponse, error) {
	var resp btrip.AlitripBtripCostCenterNewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
