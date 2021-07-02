package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCostCenterEntitySet 设置外部成本中心人员信息
// alitrip.btrip.cost.center.entity.set
//
// 设置外部成本中心人员信息
func AlitripBtripCostCenterEntitySet(clt *core.SDKClient, req *btrip.AlitripBtripCostCenterEntitySetAPIRequest, session string) (*btrip.AlitripBtripCostCenterEntitySetAPIResponse, error) {
	var resp btrip.AlitripBtripCostCenterEntitySetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
