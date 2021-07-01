package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

/* AlitripBtripCostCenterEntityAdd
增加外部成本中心人员信息
alitrip.btrip.cost.center.entity.add

增加外部成本中心人员信息 */
func AlitripBtripCostCenterEntityAdd(clt *core.SDKClient, req *btrip.AlitripBtripCostCenterEntityAddAPIRequest, session string) (*btrip.AlitripBtripCostCenterEntityAddAPIResponse, error) {
	var resp btrip.AlitripBtripCostCenterEntityAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
