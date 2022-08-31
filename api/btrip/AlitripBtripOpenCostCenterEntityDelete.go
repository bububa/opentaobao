package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterEntityDelete 删除成本中心人员信息
// alitrip.btrip.open.cost.center.entity.delete
//
// 删除成本中心人员信息
func AlitripBtripOpenCostCenterEntityDelete(clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterEntityDeleteAPIRequest, session string) (*btrip.AlitripBtripOpenCostCenterEntityDeleteAPIResponse, error) {
	var resp btrip.AlitripBtripOpenCostCenterEntityDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
