package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterEntitySet 设置成本中心人员信息
// alitrip.btrip.open.cost.center.entity.set
//
// 设置成本中心人员信息
func AlitripBtripOpenCostCenterEntitySet(clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterEntitySetAPIRequest, session string) (*btrip.AlitripBtripOpenCostCenterEntitySetAPIResponse, error) {
	var resp btrip.AlitripBtripOpenCostCenterEntitySetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
